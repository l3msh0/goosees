package main

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"errors"
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"os"
	"path/filepath"
)

// Load config from specified YAML file.
func loadConfig() (confs []*goose.DBConf, err error) {

	// migrationsDir should be placed at conf directory
	migrationsDir := filepath.Dir(prmConfPath) + "/" + prmGroup

	f, err := yaml.ReadFile(prmConfPath)
	if err != nil {
		return nil, err
	}

	rootMap, err := nodeToMap(f.Root)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("conf file syntax is not valid"))
	}

	confNodes, err := nodeToList(rootMap[prmGroup])
	if err != nil {
		return nil, errors.New(fmt.Sprintf(`group "%s" is not found on %s`, prmGroup, prmConfPath))
	}

	confs = make([]*goose.DBConf, 0, confNodes.Len())
	for i, confNode := range confNodes {
		conf, err := nodeToMap(confNode)
		if err != nil {
			return nil, err
		}

		drvNode, err := yaml.Child(conf, "driver")
		if err != nil {
			return nil, errors.New(fmt.Sprintf(`%s[%d]: required item "driver" is missing `, prmGroup, i))
		}
		drv, err := nodeToString(drvNode)
		if err != nil {
			return nil, err
		}
		drv = os.ExpandEnv(drv)

		openNode, err := yaml.Child(conf, "open")
		if err != nil {
			return nil, errors.New(fmt.Sprintf(`%s[%d]: required item "open" is missing`, prmGroup, i))
		}
		open, err := nodeToString(openNode)
		if err != nil {
			return nil, err
		}
		open = os.ExpandEnv(open)

		d, err := newDBDriver(drv, open)
		if err != nil {
			return nil, err
		}

		if !d.IsValid() {
			return nil, errors.New(fmt.Sprintf("Invalid DBConf: %v", d))
		}

		confs = append(confs, &goose.DBConf{
			MigrationsDir: migrationsDir,
			Env:           fmt.Sprintf("%s[%d]", prmGroup, i),
			Driver:        d,
		})
	}
	return confs, nil
}

// Create a new DBDriver and populate driver specific
// fields for drivers that we know about.
// Further customization may be done in NewDBConf
func newDBDriver(name, open string) (goose.DBDriver, error) {

	d := goose.DBDriver{
		Name:    name,
		OpenStr: open,
	}

	// currently only support mysql
	switch name {
	case "mysql":
		d.Import = "github.com/go-sql-driver/mysql"
		d.Dialect = &goose.MySqlDialect{}
		return d, nil
	}
	return d, errors.New("driver name is not valid")
}

// Cast YAML Node to Map
func nodeToMap(node yaml.Node) (yaml.Map, error) {
	m, ok := node.(yaml.Map)
	if !ok {
		return nil, errors.New(fmt.Sprintf("%v is not of type map", node))
	}
	return m, nil
}

// Cast YAML Node to List
func nodeToList(node yaml.Node) (yaml.List, error) {
	m, ok := node.(yaml.List)
	if !ok {
		return nil, errors.New(fmt.Sprintf("%v is not of type list", node))
	}
	return m, nil
}

// Cast YAML Node to string
func nodeToString(node yaml.Node) (string, error) {
	m, ok := node.(yaml.Scalar)
	if !ok {
		return "", errors.New(fmt.Sprintf("%v is not of type scalar", node))
	}
	return m.String(), nil
}
