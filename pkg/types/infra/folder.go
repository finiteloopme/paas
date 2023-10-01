package infra

type folderType struct {
	Folder *ResourceType `yaml:"folder" json:"folder"`
}

func (res *ResourceType) NewFolderInOrgType() {
	folder := folderType{
		Folder: &ResourceType{
			Parent: ParentType{
				Kind: "organisation",
			},
			Kind: "folderInOrganisation",
		},
	}

	res = folder.Folder
}
