package structs

import orderedmap "github.com/wk8/go-ordered-map/v2"

// ----------------------- Struct for articles.yml ----------------------- //
type Article struct {
	Emoji string `yaml:"emoji"`
	Date  string `yaml:"date"`
	Lang  string `yaml:"lang"`
	Title string `yaml:"title"`
}

// ----------------------- Struct for links.yml ----------------------- //
type Links struct {
	Wikipedia     orderedmap.OrderedMap[string, string] `yaml:"wikipedia"`
	Wiktionary    orderedmap.OrderedMap[string, string] `yaml:"wiktionary"`
	OtherSites    orderedmap.OrderedMap[string, string] `yaml:"other_sites"`
	PeopleIAdmire orderedmap.OrderedMap[string, string] `yaml:"people_i_admire"`
}

// ----------------------- Struct for languages.yml ----------------------- //
type Languages struct {
	En map[string]string `yaml:"en"`
	Es map[string]string `yaml:"es"`
}

// ----------------------- Struct for projects.yml ----------------------- //

type Project struct {
	En ProjectLanguage `yaml:"en"`
	Es ProjectLanguage `yaml:"es"`
}

type ProjectLanguage struct {
	Projects orderedmap.OrderedMap[string, ProjectDetail] `yaml:"projects"`
}

type ProjectDetail struct {
	Name        string `yaml:"name"`
	Emoji       string `yaml:"emoji"`
	Description string `yaml:"description"`
	Url         string `yaml:"url"`
	Repo        string `yaml:"repo"`
}

// ----------------------- Struct for curriculum.yml ----------------------- //
type CV struct {
	En Language `yaml:"en"`
	Es Language `yaml:"es"`
}

type Language struct {
	Headings  Headings   `yaml:"headings"`
	Education Section    `yaml:"education"`
	Work      Section    `yaml:"work"`
	Other     OtherItems `yaml:"other"`
}

type Headings struct {
	AboutMe string `yaml:"aboutme"`
	CvTitle string `yaml:"cvtitle"`
}

type Section struct {
	Heading string   `yaml:"heading"`
	List    []Detail `yaml:"list"`
}

type Detail struct {
	Title       string `yaml:"title"`
	Institution string `yaml:"institution"`
	Dates       string `yaml:"dates"`
	Location    string `yaml:"location"`
	Schedule    string `yaml:"schedule"`
	Detail      string `yaml:"detail"`
}

type OtherItems struct {
	SkillsLangs SkillsLangs `yaml:"languages"`
	Skills      Skills      `yaml:"skills"`
}

type SkillsLangs struct {
	Heading string        `yaml:"heading"`
	Level   LanguageLevel `yaml:"level"`
	List    LanguageList  `yaml:"list"`
}

type LanguageLevel struct {
	Native       string `yaml:"native"`
	Intermediate string `yaml:"intermediate"`
	Basic        string `yaml:"basic"`
}

type LanguageList struct {
	Native       []string `yaml:"native"`
	Intermediate []string `yaml:"intermediate"`
	Basic        []string `yaml:"basic"`
}

type Skills struct {
	Heading          string         `yaml:"heading"`
	SkillsCategories SkillsCategory `yaml:"skills_categories"`
	List             SkillsList     `yaml:"list"`
}

type SkillsCategory struct {
	Languages string `yaml:"languages"`
	Data      string `yaml:"data"`
	Tools1    string `yaml:"tools1"`
	Tools2    string `yaml:"tools2"`
	Tools3    string `yaml:"tools3"`
	Other1    string `yaml:"other1"`
	Other2    string `yaml:"other2"`
}

type SkillsList struct {
	NoLevel SkillsDetail `yaml:"noLevel"`
}

type SkillsDetail struct {
	Languages string `yaml:"languages"`
	Data      string `yaml:"data"`
	Tools1    string `yaml:"tools1"`
	Tools2    string `yaml:"tools2"`
	Tools3    string `yaml:"tools3"`
	Other1    string `yaml:"other1"`
	Other2    string `yaml:"other2"`
}
