package main

import (
//	"encoding/json"
	"fmt"
	"strconv"
//	"path/filepath"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/infrakit/pkg/spi"
	"github.com/docker/infrakit/pkg/spi/instance"
	"github.com/docker/infrakit/pkg/types"
        "github.com/softlayer/softlayer-go/datatypes"
        "github.com/softlayer/softlayer-go/services"
        "github.com/softlayer/softlayer-go/session"
        "github.com/softlayer/softlayer-go/sl"
//	"github.com/spf13/afero"
//	"math/rand"
)

// This example uses local files as a representation of an instance.  When we
// create an instance, we write a file in a directory.  The content of the file is simply
// the message in the provision spec, so we can verify correctness of the content easily.
// When we destroy an instance, we remove the file.
// DescribeInstances simply would list the files with the matching
// tags.

// Spec is just whatever that can be unmarshalled into a generic JSON map
type Spec map[string]map[string]string

/*
func init() {
//	rand.Seed(time.Now().UTC().UnixNano())

// sess init in NewFileInstancePlugin() or here?
}
*/

// fileInstance represents a single file instance on disk.
// change name to SLInstance or softlayerInstance
type fileInstance struct {
	instance.Description
//type Description struct {
//	ID        ID
//	LogicalID *LogicalID
//	Tags      map[string]string
//}
	Spec instance.Spec
// Spec is a specification of an instance to be provisioned
//type Spec struct {
//	// Properties is the opaque instance plugin configuration.
//	Properties *types.Any
//	// Tags are metadata that describes an instance.
//	Tags map[string]string
//	// Init is the boot script to execute when the instance is created.
//	Init string
//	// LogicalID is the logical identifier assigned to this instance, which may be absent.
//	LogicalID *LogicalID
//	// Attachments are instructions for external entities that should be attached to the instance.
//	Attachments []Attachment
//}
}


// change name to softlayerPlugin
type plugin struct {
//	name(evtually name of SL acc) or id
	Name string
//	Dir string
//	service
	sess *session.Session
//	fs  afero.Fs // is an interface for working with Files
}

/*
// NewFileInstancePlugin returns an instance plugin backed by disk files.
func NewFileInstancePlugin(dir string) instance.Plugin {
	log.Debugln("file instance plugin. dir=", dir)
	return &plugin{
		Dir: dir,
		fs:  afero.NewOsFs(), // returns a pointer to a class struct
				// is a function or a *struct returned ??
	}
}
*/

// change name to NewSLInstancePlugin or NewSoftlayerInstancePlugin
func NewFileInstancePlugin(name string) instance.Plugin {
	log.Debugln("softlayer instance plugin of name =", name)
	sess := session.New() // default endpoint https://api.softlayer.com/rest/v3/...json
	sess.Debug = true

	return &plugin{
		Name: name,
		sess: sess,
				// a struct is returned ??
	}
}


// Info returns a vendor specific name and version
// what is this for??
func (p *plugin) VendorInfo() *spi.VendorInfo {
	return &spi.VendorInfo{
		InterfaceSpec: spi.InterfaceSpec{
			Name:    "infrakit-instance-sl",
			Version: "0.3.0",
		},
		URL: "https://github.com/docker/infrakit",
	}
}

// ExampleProperties returns the properties / config of this plugin
//				where is this methode used???
func (p *plugin) ExampleProperties(/*	spec instance.Spec*/) *types.Any {

/*	if spec.Properties == nil {
		return nil, errors.New("Properties must be set")
	}*/

// Instance input, data that the instance uses to be prvisioned (request data)
/*        vGuestTemplate := datatypes.Virtual_Guest{
                Hostname:                     sl.String("mphauto"),
                Domain:                       sl.String("mphautobusiness.com"),
                MaxMemory:                    sl.Int(1024),
                StartCpus:                    sl.Int(1),
                Datacenter:                   &datatypes.Location{Name: sl.String("fra02")},
                OperatingSystemReferenceCode: sl.String("UBUNTU_LATEST"),
                LocalDiskFlag:                sl.Bool(true),
                HourlyBillingFlag:            sl.Bool(true),
        }
*/
	any, err := types.AnyValue(Spec{
		"exampleString": "a_string",
		"exampleBool":   true,
		"exampleInt":    1,
	})
	if err != nil {
		return nil
	}
	return any
}

type SpecPropertiesFormat struct {
	Type  string                 `json:"type"`
	Value map[string]interface{} `json:"value"`
}

// Validate performs local validation on a provision request.
func (p *plugin) Validate(req *types.Any) error {
	log.Debugln("validate", req.String())

	spec := Spec{}
	if err := req.Decode(&spec); err != nil {
		return err
	}

// is json(Spec) correct
//	spec.hostname == nill {
//		return someerror
//	}
/*
 24         m := map[string]Book{"foo": Book{Title: "Bar"}, "quux": Book{Title: "Baz"}}
 25         v := reflect.ValueOf(m)
 26         if v.Kind() == reflect.Map {
 27                 for _, key := range v.MapKeys() {
 28                         strct := v.MapIndex(key)
 29                         fmt.Println(key.Interface(), strct.Interface())
 30                 }
 31         }
*/
/*
	var properties interface{} = spec["Properties"]
	m := map[string]string{}
	switch p := properties.(type) {
	case map[string]interface{}:
		for k, v := range p {
			log.Debugln(k, ": ", v)
			m[k] = v
		}
	default:
		log.Debugf("Failed to get type %T", p)
	}
	mm := map[string]string{}
	switch pp := m.(type) {
	case map[string]string{}:
		for k, v := range m {
			log.Debugln(k, ":: ", v)
			mm[k] = v
		}
	default:
		log.Debugf("2failed Type=%T\n", pp)
	}
*/


	log.Debugln("Validated:", spec)
	return nil
}

// Provision creates a new instance based on the spec.
func (p *plugin) Provision(spec instance.Spec) (*instance.ID, error) {

/*
	convert spec => vGuestTemplate

// hardcoded version
*/

/*
        vGuestTemplate := datatypes.Virtual_Guest{
                Hostname:                     sl.String(spec.Properties.Hostname),
                Domain:                       sl.String(spec.Properties.Domain),
                MaxMemory:                    sl.Int(spec.Properties.MaxMemory),
                StartCpus:                    sl.Int(spec.Properties.StartCpus),
                Datacenter:                   &datatypes.Location{Name: sl.String(spec.Properties.Datacenter)},
                OperatingSystemReferenceCode: sl.String(spec.Properties.OperatingSystemReferenceCode),
                LocalDiskFlag:                sl.Bool(spec.Properties.LocalDiskFlag),
                HourlyBillingFlag:            sl.Bool(spec.Properties.HourlyBillingFlag),
        }
*/

        vGuestTemplate := datatypes.Virtual_Guest{
                Hostname:                     sl.String("mphauto"),
                Domain:                       sl.String("mphautobusiness.com"),
                MaxMemory:                    sl.Int(1024),
                StartCpus:                    sl.Int(1),
                Datacenter:                   &datatypes.Location{Name: sl.String("fra02")},
                OperatingSystemReferenceCode: sl.String("UBUNTU_LATEST"),
                LocalDiskFlag:                sl.Bool(true),
                HourlyBillingFlag:            sl.Bool(true),
        }


	service := services.GetVirtualGuestService(p.sess)
        vGuest, err := service.Mask("id;domain").CreateObject(&vGuestTemplate)
	id := instance.ID(string(*vGuest.Id))
        if err != nil {
                fmt.Printf("%s\n", err)
                return nil, err
        } else {
                fmt.Printf("\nNew Virtual Guest created with ID %d\n", id)
                fmt.Printf("Domain: %s\n", *vGuest.Domain)
        }

/*
	// simply writes a file
	// use timestamp as instance id
	id := instance.ID(fmt.Sprintf("instance-%d", rand.Int63()))
	buff, err := json.MarshalIndent(fileInstance{
		Description: instance.Description{
			Tags:      spec.Tags,
			ID:        id,
			LogicalID: spec.LogicalID,
		},
		Spec: spec,
	}, "", "")
	log.Debugln("provision", id, "data=", string(buff), "err=", err)
	if err != nil {
		return nil, err
	}
	return &id, afero.WriteFile(p.fs, filepath.Join(p.Dir, string(id)), buff, 0644)
*/

	log.Debugln("provisio", id, "data=empty")
	return &id, nil
}

// Label labels the instance
func (p *plugin) Label(instance instance.ID, labels map[string]string) error {
/*
	fp := filepath.Join(p.Dir, string(instance))
	buff, err := afero.ReadFile(p.fs, fp)
	if err != nil {
		return err
	}
	instanceData := fileInstance{}
	err = json.Unmarshal(buff, &instanceData)
	if err != nil {
		return err
	}

	if instanceData.Description.Tags == nil {
		instanceData.Description.Tags = map[string]string{}
	}
	for k, v := range labels {
		instanceData.Description.Tags[k] = v
	}

	buff, err = json.MarshalIndent(instanceData, "", "")
	log.Debugln("label:", instance, "data=", string(buff), "err=", err)
	if err != nil {
		return err
	}
	return afero.WriteFile(p.fs, fp, buff, 0644)
*/
	return nil
}

// Destroy terminates an existing instance.
func (p *plugin) Destroy(instance instance.ID) error {
        fmt.Println("Deleting virtual guest")

	service := services.GetVirtualGuestService(p.sess)

        // Wait for transactions to finish
        fmt.Printf("Waiting for transactions to complete before destroying.")
	id, err := strconv.Atoi(string(instance))
	if err != nil {
		log.Debug("Atoi conversion error", err)
		return err
	}
        service = service.Id(id)

        // Delay to allow transactions to be registered
        time.Sleep(10 * time.Second)
// what does service.GetActiveTransactions() do?
        for transactions, _ := service.GetActiveTransactions(); len(transactions) > 0; {
                fmt.Print(".")
                time.Sleep(10 * time.Second)
                transactions, err = service.GetActiveTransactions()
        }


//	service := services.GetVirtualGuestService(p.sess)

        success, err := service.DeleteObject()
        if err != nil {
                fmt.Printf("Error deleting virtual guest: %s", err)
        } else if success == false {
                fmt.Printf("Error deleting virtual guest")
        } else {
                fmt.Printf("Virtual Guest deleted successfully")
        }

	log.Debugln("destroy", instance)
	return err

/*
	fp := filepath.Join(p.Dir, string(instance))
	log.Debugln("destroy", fp)
	return p.fs.Remove(fp)
*/
}

// DescribeInstances returns descriptions of all instances matching all of the provided tags.
// TODO - need to define the fitlering of tags => AND or OR of matches?
func (p *plugin) DescribeInstances(tags map[string]string) ([]instance.Description, error) {
/*
	log.Debugln("describe-instances", tags)
	entries, err := afero.ReadDir(p.fs, p.Dir)
	if err != nil {
		return nil, err
	}

	result := []instance.Description{}
scan:
	for _, entry := range entries {
		fp := filepath.Join(p.Dir, entry.Name())
		file, err := p.fs.Open(fp)
		if err != nil {
			log.Warningln("error opening", fp)
			continue scan
		}

		inst := fileInstance{}
		err = json.NewDecoder(file).Decode(&inst)
		if err != nil {
			log.Warning("cannot decode", entry.Name())
			continue scan
		}

		if len(tags) == 0 {
			result = append(result, inst.Description)
		} else {
			for k, v := range tags {
				if inst.Tags[k] != v {
					continue scan // we implement AND
				}
			}
			result = append(result, inst.Description)
		}

	}
	return result, nil
*/
	return nil, nil
}
