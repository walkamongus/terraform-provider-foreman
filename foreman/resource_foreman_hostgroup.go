package foreman

import (
	"fmt"
	"strconv"

	"github.com/wayfair/terraform-provider-foreman/foreman/api"
	"github.com/wayfair/terraform-provider-utils/autodoc"
	"github.com/wayfair/terraform-provider-utils/log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceForemanHostgroup() *schema.Resource {
	return &schema.Resource{

		Create: resourceForemanHostgroupCreate,
		Read:   resourceForemanHostgroupRead,
		Update: resourceForemanHostgroupUpdate,
		Delete: resourceForemanHostgroupDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

			autodoc.MetaAttribute: &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
				Description: fmt.Sprintf(
					"%s Hostgroups are organized in a tree-like structure and inherit "+
						"values from their parent hostgroup(s). When hosts get associated "+
						"with a hostgroup, it will inherit attributes from the hostgroup. "+
						"This allows for easy, shared configuration of various hosts based "+
						"on common attributes.",
					autodoc.MetaSummary,
				),
			},

			"title": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Description: "The title is the fullname of the hostgroup.  A " +
					"hostgroup's title is a path-like string from the head " +
					"of the hostgroup tree down to this hostgroup.  The title will be " +
					"in the form of: \"<parent 1>/<parent 2>/.../<name>\".",
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Description: fmt.Sprintf(
					"Hostgroup name. "+
						"%s \"compute\"",
					autodoc.MetaExample,
				),
			},

			// -- Foreign Key Relationships --

			"architecture_id": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntAtLeast(0),
				Description:  "ID of the architecture associated with this hostgroup.",
			},

			"compute_profile_id": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntAtLeast(0),
				Description:  "ID of the compute profile associated with this hostgroup.",
			},

			"domain_id": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntAtLeast(0),
				Description:  "ID of the domain associated with this hostgroup.",
			},

			"environment_id": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntAtLeast(0),
				Description:  "ID of the environment associated with this hostgroup.",
			},

			"medium_id": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntAtLeast(0),
				Description:  "ID of the media associated with this hostgroup.",
			},

			"operatingsystem_id": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntAtLeast(0),
				Description:  "ID of the operating system associated with this hostgroup.",
			},

			"parameters": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "ID of the operating system associated with this hostgroup.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
						"created_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"updated_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"parent_id": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntAtLeast(0),
				Description:  "ID of the parent hostgroup.",
			},

			"ptable_id": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntAtLeast(0),
				Description:  "ID of the partition table associated with this hostgroup.",
			},

			"puppet_ca_proxy_id": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntAtLeast(0),
				Description: "ID of the smart proxy acting as the puppet certificate " +
					"authority server for this hostgroup.",
			},

			"puppet_proxy_id": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntAtLeast(0),
				Description: "ID of the smart proxy acting as the puppet proxy " +
					"server for this hostgroup.",
			},

			"realm_id": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntAtLeast(0),
				Description:  "ID of the realm associated with this hostgroup.",
			},

			"subnet_id": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntAtLeast(0),
				Description:  "ID of the subnet associated with the hostgroup.",
			},
		},
	}
}

// -----------------------------------------------------------------------------
// Conversion Helpers
// -----------------------------------------------------------------------------

// buildForemanHostgroup constructs a ForemanHostgroup struct from a resource
// data reference. The struct's members are populated from the data populated
// in the resource data. Missing members will be left to the zero value for
// that member's type.
func buildForemanHostgroup(d *schema.ResourceData) *api.ForemanHostgroup {
	log.Tracef("resource_foreman_hostgroup.go#buildForemanHostgroup")

	hostgroup := api.ForemanHostgroup{}

	obj := buildForemanObject(d)
	hostgroup.ForemanObject = *obj

	var attr interface{}
	var ok bool

	if attr, ok = d.GetOk("title"); ok {
		hostgroup.Title = attr.(string)
	}

	if attr, ok = d.GetOk("architecture_id"); ok {
		hostgroup.ArchitectureId = attr.(int)
	}

	if attr, ok = d.GetOk("compute_profile_id"); ok {
		hostgroup.ComputeProfileId = attr.(int)
	}

	if attr, ok = d.GetOk("domain_id"); ok {
		hostgroup.DomainId = attr.(int)
	}

	if attr, ok = d.GetOk("environment_id"); ok {
		hostgroup.EnvironmentId = attr.(int)
	}

	if attr, ok = d.GetOk("medium_id"); ok {
		hostgroup.MediaId = attr.(int)
	}

	if attr, ok = d.GetOk("operatingsystem_id"); ok {
		hostgroup.OperatingSystemId = attr.(int)
	}

	// I don't know what black magic needs to happen for running tests
	// uncommenting the following will cause tests to fail with type errors
	//if attr, ok = d.GetOk("parameters.#"); ok {
	//	params := make([]api.ForemanHostgroupParameter, attr.(int))
	//	for i := 0; i < attr.(int); i++ {
	//		idx := strconv.Itoa(i)
	//		param := api.ForemanHostgroupParameter{
	//			ID:        d.Get("parameters." + idx + "id").(int64),
	//			Name:      d.Get("parameters." + idx + "name").(string),
	//			Value:     d.Get("parameters." + idx + "value").(string),
	//			CreatedAt: d.Get("parameters." + idx + "created_at").(string),
	//			UpdatedAt: d.Get("parameters." + idx + "updated_at").(string),
	//			Priority:  d.Get("parameters." + idx + "priority").(int64),
	//		}
	//		params = append(params, param)
	//	}
	//	hostgroup.Parameters = params
	//}

	if attr, ok = d.GetOk("parent_id"); ok {
		hostgroup.ParentId = attr.(int)
	}

	if attr, ok = d.GetOk("ptable_id"); ok {
		hostgroup.PartitionTableId = attr.(int)
	}

	if attr, ok = d.GetOk("puppet_ca_proxy_id"); ok {
		hostgroup.PuppetCAProxyId = attr.(int)
	}

	if attr, ok = d.GetOk("puppet_proxy_id"); ok {
		hostgroup.PuppetProxyId = attr.(int)
	}

	if attr, ok = d.GetOk("realm_id"); ok {
		hostgroup.RealmId = attr.(int)
	}

	if attr, ok = d.GetOk("subnet_id"); ok {
		hostgroup.SubnetId = attr.(int)
	}

	return &hostgroup
}

// setResourceDataFromForemanHostgroup sets a ResourceData's attributes from
// the attributes of the supplied ForemanHostgroup struct
func setResourceDataFromForemanHostgroup(d *schema.ResourceData, fh *api.ForemanHostgroup) {
	log.Tracef("resource_foreman_hostgroup.go#setResourceDataFromForemanHostgroup")

	params := make([]map[string]interface{}, len(fh.Parameters))
	for k, v := range fh.Parameters {
		param := make(map[string]interface{})
		param["name"] = v.Name
		param["value"] = v.Value
		param["created_at"] = v.CreatedAt
		param["id"] = v.ID
		param["priority"] = v.Priority
		param["updated_at"] = v.UpdatedAt
		params[k] = param
	}

	d.SetId(strconv.Itoa(fh.Id))

	if err := d.Set("title", fh.Title); err != nil {
		log.Errorf("error setting hostgroup title: %s", err)
	}
	if err := d.Set("name", fh.Name); err != nil {
		log.Errorf("error setting hostgroup name: %s", err)
	}
	if err := d.Set("architecture_id", fh.ArchitectureId); err != nil {
		log.Errorf("error setting hostgroup architecture_id: %s", err)
	}
	if err := d.Set("compute_profile_id", fh.ComputeProfileId); err != nil {
		log.Errorf("error setting hostgroup compute_profile_id: %s", err)
	}
	if err := d.Set("domain_id", fh.DomainId); err != nil {
		log.Errorf("error setting hostgroup domain_id: %s", err)
	}
	if err := d.Set("environment_id", fh.EnvironmentId); err != nil {
		log.Errorf("error setting hostgroup environment_id: %s", err)
	}
	if err := d.Set("medium_id", fh.MediaId); err != nil {
		log.Errorf("error setting hostgroup medium_id: %s", err)
	}
	if err := d.Set("operatingsystem_id", fh.OperatingSystemId); err != nil {
		log.Errorf("error setting hostgroup operatingsystem_id: %s", err)
	}
	if err := d.Set("parameters", params); err != nil {
		log.Errorf("error setting hostgroup parameters: %s", err)
	}
	if err := d.Set("parent_id", fh.ParentId); err != nil {
		log.Errorf("error setting hostgroup parent_id: %s", err)
	}
	if err := d.Set("ptable_id", fh.PartitionTableId); err != nil {
		log.Errorf("error setting hostgroup ptable_id: %s", err)
	}
	if err := d.Set("puppet_ca_proxy_id", fh.PuppetCAProxyId); err != nil {
		log.Errorf("error setting hostgroup puppet_ca_proxy_id: %s", err)
	}
	if err := d.Set("puppet_proxy_id", fh.PuppetProxyId); err != nil {
		log.Errorf("error setting hostgroup puppet_proxy_id: %s", err)
	}
	if err := d.Set("realm_id", fh.RealmId); err != nil {
		log.Errorf("error setting hostgroup realm_id: %s", err)
	}
	if err := d.Set("subnet_id", fh.SubnetId); err != nil {
		log.Errorf("error setting hostgroup subnet_id: %s", err)
	}
}

// -----------------------------------------------------------------------------
// Resource CRUD Operations
// -----------------------------------------------------------------------------

func resourceForemanHostgroupCreate(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_hostgroup.go#Create")

	client := meta.(*api.Client)
	h := buildForemanHostgroup(d)

	log.Debugf("ForemanHostgroup: [%+v]", h)

	createdHostgroup, createErr := client.CreateHostgroup(h)
	if createErr != nil {
		return createErr
	}

	log.Debugf("Created ForemanHostgroup: [%+v]", createdHostgroup)

	setResourceDataFromForemanHostgroup(d, createdHostgroup)

	return nil
}

func resourceForemanHostgroupRead(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_hostgroup.go#Read")

	client := meta.(*api.Client)
	h := buildForemanHostgroup(d)

	log.Debugf("ForemanHostgroup: [%+v]", h)

	readHostgroup, readErr := client.ReadHostgroup(h.Id)
	if readErr != nil {
		return readErr
	}

	log.Debugf("Read ForemanHostgroup: [%+v]", readHostgroup)

	setResourceDataFromForemanHostgroup(d, readHostgroup)

	return nil
}

func resourceForemanHostgroupUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_hostgroup.go#Update")

	// TODO(ALL): 404 errors here (for v.1.11.4 ) - i think we need to
	//   concatentate the id with the title, replacing forward slash with a dash?
	//   getting weird behavior when updating a hostgroup aside from updating the
	//   hostgroup's name

	client := meta.(*api.Client)
	h := buildForemanHostgroup(d)

	log.Debugf("ForemanHostgroup: [%+v]", h)

	updatedHostgroup, updateErr := client.UpdateHostgroup(h)
	if updateErr != nil {
		return updateErr
	}

	log.Debugf("Updated ForemanHostgroup: [%+v]", updatedHostgroup)

	setResourceDataFromForemanHostgroup(d, updatedHostgroup)

	return nil
}

func resourceForemanHostgroupDelete(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_hostgroup.go#Delete")

	client := meta.(*api.Client)
	h := buildForemanHostgroup(d)

	log.Debugf("ForemanHostgroup: [%+v]", h)

	// NOTE(ALL): d.SetId("") is automatically called by terraform assuming delete
	//   returns no errors
	return client.DeleteHostgroup(h.Id)
}
