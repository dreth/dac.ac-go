# importing libs
import os
import re
from time import sleep

from ruamel.yaml import YAML

# Configure ruamel.yaml instance for ASCII output
yaml = YAML()
yaml.encoding = "utf-8"
yaml.allow_unicode = False


# %%
# emoji patterns remove emojis
# https://stackoverflow.com/a/49146722/330558
def remove_emoji(string):
    emoji_pattern = re.compile(
        "["
        "\U0001F600-\U0001F64F"  # emoticons
        "\U0001F300-\U0001F5FF"  # symbols & pictographs
        "\U0001F680-\U0001F6FF"  # transport & map symbols
        "\U0001F1E0-\U0001F1FF"  # flags (iOS)
        "\U00002500-\U00002BEF"  # chinese char
        "\U00002702-\U000027B0"
        "\U00002702-\U000027B0"
        "\U000024C2-\U0001F251"
        "\U0001f926-\U0001f937"
        "\U00010000-\U0010ffff"
        "\u2640-\u2642"
        "\u2600-\u2B55"
        "\u200d"
        "\u23cf"
        "\u23e9"
        "\u231a"
        "\ufe0f"  # dingbats
        "\u3030"
        "]+\s",
        flags=re.UNICODE,
    )
    return emoji_pattern.sub(r"", string)


# remove list items in detail sections
def remove_lists(string):
    item_pattern = re.compile("-\s")
    return item_pattern.sub(r"", string)


# escape som characters
def escape_chars(string):
    chars = ["&"]
    # replace stuff
    for char in chars:
        string = string.replace(char, f"\\{char}")
    return string


# %%
# CV data
cv_data = {"en": {}, "es": {}}

# basic links
basic_links = {
    "website": {"link": "https://dac.ac", "name": "dac.ac"},
    "github": {"link": "https://github.com/dreth", "name": "@dreth"},
    "email": {"link": "mailto:daniel@m.dac.ac", "name": "d@dac.ac"},
    "linkedin": {"link": "https://linkedin.com/in/dreth", "name": "@dreth"},
}

# other cv data
with open("./data/curriculum.yml", "r") as f:
    full_cv_yaml = yaml.load(remove_emoji(f.read()))

    for lang in cv_data.keys():
        cv_data[lang]["education"] = dict(full_cv_yaml[lang]["education"])
        cv_data[lang]["work"] = dict(full_cv_yaml[lang]["work"])
        cv_data[lang]["languages"] = dict(full_cv_yaml[lang]["other"]["languages"])
        cv_data[lang]["skills"] = dict(full_cv_yaml[lang]["other"]["skills"])

    # remove headings
    for lang in cv_data.keys():
        for section in ["education", "work", "languages", "skills"]:
            del cv_data[lang][section]["heading"]

    # remove dashes in detail section
    detailed_sections = ["work", "education"]
    for lang in cv_data.keys():
        for section in detailed_sections:
            for i, element_list in enumerate(cv_data[lang][section]["list"]):
                cv_data[lang][section]["list"][i]["detail"] = remove_lists(
                    element_list["detail"]
                )

# open lang data
with open("./data/languages.yml", "r") as f:
    lang_data = yaml.load(f)

# %%
# skeleton
cv_skeleton = r"""
%------------------------
% Resume Template
% Author : Anubhav Singh
% Github : https://github.com/xprilion
% License : MIT
%------------------------

\documentclass[a4paper,20pt]{article}

\usepackage[inline]{enumitem}
\usepackage{latexsym}
\usepackage[empty]{fullpage}
\usepackage{titlesec}
\usepackage{marvosym}
\usepackage[usenames,dvipsnames]{color}
\usepackage{verbatim}
\usepackage{enumitem}
\usepackage[pdftex]{hyperref}
\usepackage{fancyhdr}

\pagestyle{fancy}
\fancyhf{} % clear all header and footer fields
\fancyfoot{}
\renewcommand{\headrulewidth}{0pt}
\renewcommand{\footrulewidth}{0pt}

% Adjust margins
\addtolength{\oddsidemargin}{-0.530in}
\addtolength{\evensidemargin}{-0.375in}
\addtolength{\textwidth}{1in}
\addtolength{\topmargin}{-0.7in}
\addtolength{\textheight}{1.11in}

\urlstyle{rm}

\raggedbottom
\raggedright
\setlength{\tabcolsep}{0in}

% Sections formatting
\titleformat{\section}{
  \vspace{-10pt}\scshape\raggedright\large
}{}{0em}{}[\color{black}\titlerule \vspace{-6pt}]

%-------------------------
% Custom commands
\newcommand{\resumeItem}[2]{
  \item\small{
    \textbf{#1}{: #2 \vspace{-2pt}}
  }
}

\newcommand{\resumeItemWithoutTitle}[1]{
  \item\small{
    {\vspace{-2pt}}
  }
}

\newcommand{\resumeSubheading}[4]{
  \vspace{-1pt}\item
    \begin{tabular*}{0.97\textwidth}{l@{\extracolsep{\fill}}r}
      \textbf{#1} & #2 \\
      \textit{#3} & \textit{#4} \\
    \end{tabular*}\vspace{-5pt}
}


\newcommand{\resumeSubItem}[2]{\resumeItem{#1}{#2}\vspace{-3pt}}

\renewcommand{\labelitemii}{$\circ$}

\newcommand{\resumeSubHeadingListStart}{\begin{itemize}[leftmargin=*]}
\newcommand{\resumeSubHeadingListEnd}{\end{itemize}}
\newcommand{\resumeItemListStart}{\begin{itemize}}
\newcommand{\resumeItemListEnd}{\end{itemize}\vspace{-5pt}}

%-----------------------------
%%%%%%  CV STARTS HERE  %%%%%%

\begin{document}

%----------HEADING-----------------
\begin{tabular*}{\textwidth}{l@{\extracolsep{\fill}}r}
  \textbf{{\LARGE Daniel Alonso}} & Email: {\color{blue}\href{mailto:}{*email:name*}}\\
  Website: {\color{blue}\href{*website:link*}{*website:name*}} & Github: ~~~{\color{blue}\href{*github:link*}{*github:name*}}
\end{tabular*}

%----------SECTIONS TO ITERATE ON-----------------
*iteration_items*

\end{document}
"""

# %%
# blocks
cv_blocks = {
    ### SECTION NAMES PER LANG
    "section_names": {
        "en": {
            "intro": "Introduction",
            "education": "Education",
            "languages": "Languages",
            "skills": "Skills Summary",
            "work": "Work experience",
            "projects": "Blog \\& Certifications",
        },
        "es": {
            "intro": "Introducción",
            "education": "Educación",
            "languages": "Idiomas",
            "skills": "Habilidades",
            "work": "Experiencia de trabajo",
            "projects": "Blog y Certificaciones",
        },
    },
    ### GENERAL SEPARATORS
    "separator": {
        "intro": r"""
\vspace{-2pt}""",
        "education": r"""
\vspace{-2pt}""",
        "work": r"""
\vspace{-2pt}""",
        "languages": r"""
\vspace{-2pt}""",
    },
    ### GENERAL INDENTERS
    "indenter": {
        "en": {
            "skills": {
                "Languages": "~~~~~",
                "Data": "~~~~~~~~",
                "Tools": "~~~~~~~~~~~",
                "Other": "~~~~~~~~~~~~~",
            }
        },
        "es": {
            "skills": {
                "Languages": "~~~~~~",
                "Data": "~~~~~~~~",
                "Tools": "~~~~~~",
                "Other": "~~~~~~~~~~~",
            }
        },
    },
    "intro": {
        "title": r"""
%-----------INTRODUCTION-----------------
\vspace{-3pt}
\section{*intro*}
{*intro_text*}""",
        "iterative_block": "",
        "itemize_start": "",
        "itemize_item": "",
        "itemize_end": "",
        "closing_tag": "",
    },
    ### EDUCATION SECTION
    "education": {
        "title": r"""
%-----------EDUCATION-----------------
\vspace{-15pt}
\section{*education*}
\resumeSubHeadingListStart""",
        "iterative_block": r"""
  \resumeSubheading
    {*institution*}{*location*}
    {*title*}{*dates*}""",
        "itemize_start": r"""
\begin{itemize} \itemsep-0.24em""",
        "itemize_item": r"""
  \item *item*""",
        "itemize_end": r"""
\end{itemize}""",
        "closing_tag": r"""
\resumeSubHeadingListEnd""",
    },
    ### SKILLS SECTION
    "skills": {
        "title": r"""
%-----------SKILLS-----------------
\vspace{-3pt}
\section{*skills*}
  \resumeSubHeadingListStart
            """,
        "iterative_block": r"""
\vspace{-1pt}
\resumeSubItem{*section*}{*indenter**skill_list*}""",
        "closing_tag": r"""
\resumeSubHeadingListEnd""",
    },
    ### WORK SECTION
    "work": {
        "title": r"""
%-----------WORK-----------------
\vspace{-12pt}
\section{*work*}
  \resumeSubHeadingListStart""",
        "iterative_block": r"""
  \resumeSubheading{*institution*}{*location*}
    {*title* (*schedule*)}{*dates*}""",
        "itemize_start": r"""
\begin{itemize} \itemsep-0.24em""",
        "itemize_item": r"""
  \item *item*""",
        "itemize_end": r"""
\end{itemize}""",
        "closing_tag": r"""
\resumeSubHeadingListEnd""",
    },
    ### SKILLS SECTION
    "languages": {
        "title": r"""
%-----------LANGUAGES-----------------
\vspace{-15pt}
\section{*languages*}
  \begin{center}
  \begin{itemize*}""",
        "iterative_block": r"""
  \item \textbf{*level*} *language_list*""",
        "closing_tag": r"""
\end{itemize*}
\end{center}""",
    },
    ### PROJECTS SECTION
    "projects": {
        "title": r"""
%-----------PROJECTSANDBLOG-----------------
\vspace{-9pt}
\section{*projects*}
\textbf{*projects:text*}""",
        "text": {
            "en": "View all my \href{https://dac.ac/blog}{articles}, and \href{https://www.linkedin.com/in/dreth/details/certifications/}{certifications}.",
            "es": "Ver todos/as mis, \href{https://dac.ac/blog}{artículos}, y \href{https://www.linkedin.com/in/dreth/details/certifications/}{certificaciones}.",
        },
        "closing_tag": "",
    },
}


# %%
# create the CV tex document
def fill_cv(cv_skeleton, l="en"):
    # adding links to header
    for section, detail in basic_links.items():
        for subsection, value in detail.items():
            cv_skeleton = cv_skeleton.replace(f"*{section}:{subsection}*", value)

    # iteration items
    # generate iteration items string
    iteration_items = ""

    # iterating over standardized sections
    standardized_sections = [
        "intro",
        "skills",
        "languages",
        "education",
        "work",
        "projects",
    ]
    for section in standardized_sections:
        # add to the iteration items string
        iteration_items += f"""


            {cv_blocks[section]['title']}
        """

        # introuction
        if section == "intro":
            # add intro
            part1_intro = (
                "\n\item {" + lang_data[l]["cv_introduction"].split("<br><br>")[0] + "}"
            )
            part2_intro = (
                "\n\item {" + lang_data[l]["cv_introduction"].split("<br><br>")[1] + "}"
            )

            # add itemize and structure intro
            intro_text = (
                r"\begin{itemize} \itemsep-0.24em"
                + "".join([part1_intro, part2_intro])
                + "\n"
                + r"\end{itemize}"
            )

            iteration_items = iteration_items.replace(
                "*intro_text*",
                intro_text,
            )

        # items in education
        if section == "education":
            # add iterative items from specified section
            for i, entry in enumerate(cv_data[l][section]["list"]):
                # replace items in entry and append to iteration items
                it_item = cv_blocks[section]["iterative_block"]
                for key, value in entry.items():
                    # replace special tags
                    it_item = it_item.replace(f"*{key}*", value)

                # add list of items for work experience detail
                it_item += cv_blocks[section]["itemize_start"]

                # iterate over details
                for detail_item in (
                    cv_data[l][section]["list"][i]["detail"]
                    .replace("%", r"\\%")
                    .split("<br>")
                ):
                    it_item += cv_blocks[section]["itemize_item"].replace(
                        "*item*", detail_item
                    )

                # add itemize end
                it_item += cv_blocks[section]["itemize_end"]

                # add separator
                if i < (len(cv_data[l][section]["list"]) - 1):
                    it_item += cv_blocks["separator"][section]

                # add language level
                iteration_items += escape_chars(it_item)

        # items in skills section
        if section == "skills":
            # done skills section
            done_skills_sections = set()
            skills_section = ""

            # replace items in entry and append to iteration items
            for key, value in cv_data[l][section]["list"]["noLevel"].items():
                # section name
                section_name = re.sub(
                    "\d", "", cv_data[l][section]["skills_categories"][key]
                ).capitalize()
                section_name_std = re.sub("\d", "", key).capitalize()

                # replace special tags
                if section_name_std not in done_skills_sections:
                    skill_block = cv_blocks[section]["iterative_block"]

                    # tags and replacements
                    tags = ["*section*", "*indenter*", "*skill_list*"]
                    replacements = [
                        section_name,
                        cv_blocks["indenter"][l][section][section_name_std],
                        value,
                    ]

                    # make replacements
                    for tag, replacement in zip(tags, replacements):
                        skill_block = skill_block.replace(tag, replacement)

                    # add element to skills section string
                    skills_section += skill_block
                else:
                    # subtract previously added element
                    skills_section = skills_section[
                        0 : len(skills_section) - len(skill_block)
                    ]

                    # additional space in spanish
                    extra_space = {"en": "", "es": "~"}[l]

                    # adding the additional elements to the section
                    skill_block = skill_block[0:-1]
                    skill_block += f"\\\\~~~~~~~~~~~~~~~~~~~~~~~~{extra_space}{value}}}"

                    # add with the added newline
                    skills_section += skill_block

                # append to done skills sections
                done_skills_sections.add(section_name_std)

            # add skills section
            iteration_items += skills_section

        # items in languages section
        if section == "languages":
            # add languages and levels
            for i, (level, level_name) in enumerate(
                cv_data[l][section]["level"].items()
            ):
                # add replace cv block tags
                it_item = cv_blocks[section]["iterative_block"].replace(
                    "*level*", level_name
                )

                # iterate over languages in section
                it_item = it_item.replace(
                    "*language_list*",
                    "".join([f"{x}, " for x in cv_data[l][section]["list"][level]])[:-2]
                    + ".",
                )

                # add language level
                iteration_items += it_item

        # items in work section
        if section == "work":
            # add iterative items from specified section
            for i, entry in enumerate(cv_data[l][section]["list"]):
                # replace items in entry and append to iteration items
                it_item = cv_blocks[section]["iterative_block"]
                for key, value in entry.items():
                    # replace special tags
                    it_item = it_item.replace(f"*{key}*", value)

                # add list of items for work experience detail
                it_item += cv_blocks[section]["itemize_start"]

                # iterate over details
                for detail_item in (
                    cv_data[l][section]["list"][i]["detail"]
                    .replace("%", r"\\%")
                    .split("<br>")
                ):
                    it_item += cv_blocks[section]["itemize_item"].replace(
                        "*item*", detail_item
                    )

                # add itemize end
                it_item += cv_blocks[section]["itemize_end"]

                # add separator
                if i < (len(cv_data[l][section]["list"]) - 1):
                    it_item += cv_blocks["separator"][section]

                # add language level
                iteration_items += escape_chars(it_item)

        # items in the projects section
        if section == "projects":
            # add projects link
            iteration_items = iteration_items.replace(
                "*projects:text*", cv_blocks[section]["text"][l]
            )

        # add closing tag
        iteration_items += cv_blocks[section]["closing_tag"]

    # add iteration items to cv skeleton
    cv_skeleton = cv_skeleton.replace("*iteration_items*", iteration_items)
    cv_skeleton = cv_skeleton.replace(
        r"""
\begin{itemize} \itemsep-0.2em
  \item 
\end{itemize}""",
        "",
    )

    # add section name
    for section in standardized_sections + ["projects", "intro"]:
        cv_skeleton = cv_skeleton.replace(
            f"*{section}*", cv_blocks["section_names"][l][section]
        )

    # create a tex file for cv
    with open(f"./static/cv/Daniel_CV_{l}.tex", "w") as f:
        f.write(cv_skeleton)


# %%
# run function and save file corresponding to language
fill_cv(cv_skeleton=cv_skeleton, l="en")
fill_cv(cv_skeleton=cv_skeleton, l="es")

# force compile latex
os.system("cd static/cv && latexmk")

# clean up
sleep(1.5)
for filetype in ["fls", "log", "fdb_latexmk", "aux", "gz", "out"]:
    os.system(f"cd static/cv && rm *.{filetype}")
