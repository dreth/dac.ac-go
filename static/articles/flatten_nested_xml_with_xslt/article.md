# Using XSLT to flatten a nested XML

<div class="text-center pb-3 sm:pb-4">
<span class="font-black"><b>January 15, 2023</b></span></div>

XML files have been a bit annoying to deal with in my experience while working with data. However, today my friend Julio asked me for an emergency favour to help him extract data from an XML file that was *not flat*, meaning it couldn't be easily read into pandas using the pandas `read_xml()` function, which is intended to only read flat XML files.

By a flat XML file, I mean a file that looks like this:

```xml
<?xml version='1.0' encoding='utf-8'?>
<data xmlns="http://example.com">
 <row>
   <shape>square</shape>
   <degrees>360</degrees>
   <sides>4.0</sides>
 </row>
 <row>
   <shape>circle</shape>
   <degrees>360</degrees>
   <sides/>
 </row>
 <row>
   <shape>triangle</shape>
   <degrees>180</degrees>
   <sides>3.0</sides>
 </row>
</data>
```

Using pandas' `read_xml()` function without any additional parameters other than just the file or text we want to convert into a dataframe would yield:

```
>>> pd.read_xml(fl)
      shape  degrees  sides
0    square      360    4.0
1    circle      360    NaN
2  triangle      180    3.0
```

Pretty straightforward, but what happens when the file is *not flat*? Well, `read_xml()` has a parameter `stylesheet` we could pass an XSL file to in order to flatten the XML file, but we first have to create the XSL file!

***

### Article index

- [Using XSLT to flatten a nested XML](#using-xslt-to-flatten-a-nested-xml)
    - [Article index](#article-index)
  - [The task](#the-task)
  - [Flattening the XML file using XSLT](#flattening-the-xml-file-using-xslt)
    - [Connecting the XML with the XSL document](#connecting-the-xml-with-the-xsl-document)
    - [Creating the XSL file](#creating-the-xsl-file)
    - [Creating a template to flatten the file](#creating-a-template-to-flatten-the-file)
  - [Additional resources](#additional-resources)

***

## The task

OK, now we know what a *flat* XML file is and what `read_xml()` returns without any additional parameters. What should we do when the file is *not flat*? Well, by not flat I mean that not all of the contents are _at most_ 2 levels down the top node (the `<data>` tag in this case). Here's an example chunk of the file that Julio sent me:

```xml
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<sentences>
  <sentence id="2339">
    <text>I charge it at night and skip taking the cord with me because of the good battery life.</text>
    <aspectTerms>
      <aspectTerm term="cord" polarity="neutral" from="41" to="45"/>
      <aspectTerm term="battery life" polarity="positive" from="74" to="86"/>
    </aspectTerms>
  </sentence>
  <sentence id="812">
      <text>I bought a HP Pavilion DV4-1222nr laptop and have had so many problems with the computer.</text>
  </sentence>
  <sentence id="1316">
    <text>The tech guy then said the service center does not do 1-to-1 exchange and I have to direct my concern to the "sales" team, which is the retail shop which I bought my netbook from.</text>
    <aspectTerms>
      <aspectTerm term="service center" polarity="negative" from="27" to="41"/>
      <aspectTerm term="&quot;sales&quot; team" polarity="negative" from="109" to="121"/>
      <aspectTerm term="tech guy" polarity="neutral" from="4" to="12"/>
    </aspectTerms>
  </sentence>
</sentences>
```

Here, the top level node is the `<sentences>` tag, but tags under each `<sentence>` sometimes contain an `<aspectTerms>` tag which then inside contain `<aspectTerm>` tags each containing data in several of their attributes (`term`, `polarity`, `from` and `to`). Additionally, the `id` attribute of each sentence must also form part of the final dataframe. 

Therefore, as requested by Julio, the resulting dataframe had to look like this:

```
     id                                               text      aspectTerm  from    to  polarity
0  2339  I charge it at night and skip taking the cord ...            cord    41    45   neutral
1  2339  I charge it at night and skip taking the cord ...    battery life    74    86  positive
2   812  I bought a HP Pavilion DV4-1222nr laptop and h...            <NA>  <NA>  <NA>      <NA>
3  1316  The tech guy then said the service center does...  service center    27    41  negative
4  1316  The tech guy then said the service center does...    "sales" team   109   121  negative
```

Therefore, the qualities the dataframe must have are:

* The dataframe is melted so `id` and `text` are repeated *per each* element inside the `<aspectTerms>` tag inside each `<sentence>`
* The dataframe contains missing values for those sentences which have no child `<aspectTerms>` tag inside them, all those missing values *must be homogenous* (meaning no mixed numpy float `NaNs` or python `None`, ideally, all should be converted to pandas `pd.NA` missing values)
* The values in the columns `id`, `from` and `to` must be integers
* The full contents of each `<text>` tag inside each `<sentence>` tag must be contained in its specified column, same for each `term` attribute in every `<aspectTerm>` tag, including symbols like quotes and such

***

## Flattening the XML file using XSLT

In order to flatten the XML file we can look for convoluted methods where we traverse the XML file iteratively to extract information from it, but this interesting [thread on stackoverflow](https://stackoverflow.com/questions/37730436/flatten-a-nested-xml-using-xslt) prompted me to research XSLT a bit more.

XSLT stands for "Extensible Stylesheet Language Transformations", it's a language originally design for transforming XML docs into other XMLs, other markup language files like HTML, plain text, etc. It's a powerful tool that utilizes XPath to identify, extract and reference elements within the source XML file. It's an extraordinarily powerful tool to flatten XML files.

Note that I *just found out this thing exists* so bear with me if you already know everything I'm writing here, I just thought this was immensely cool and makes dealing with XML files a much less annoying endeavor.

***

### Connecting the XML with the XSL document

First, create the XSL file (with extension `.xsl`) then you must connect the XML file with its stylesheet by adding this tag at the top of the document right under the XML declaration:

```xml
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<?xml-stylesheet type="text/xsl" version="2.0" href="filename.xsl"?>
```

Where the `href` tag on the `xml-stylesheet` tag is the name of the file you want to tell the XML document to look at in order to be transformed.

***

### Creating the XSL file

XSL files are also XML documents, so they must start with the XML declaration and then you must add the following tag to reference the start of the stylesheet:

```xml
<?xml version="1.0"?>
<xsl:stylesheet version="2.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
```

Version 3.0 of XSL was released in 2017, but in this case I went with version 2.0, why? no reason whatsoever, I just copied the tag online from a tutorial on youtube I was looking at to understand XSL and XPath a little better. It's not too important in this case, as for my purposes, everything worked without issues, but it's very likely that the same code I wrote would've worked with version 3.0

The stylesheet declaration should also end with its closing tag:

```xml
</xsl:stylesheet>
```

***

### Creating a template to flatten the file

In my case, I wanted to make a template because the file contained quite a lot of data, so I started by matching the top level tag `<sentences>`:

```xml
<xsl:template match="/sentences">
</xsl:template>
```

Then, because all I wanted to do was flatten the document, I kept the same top level tag `<sentences>`:

```xml
<xsl:template match="/sentences">
  <sentences>
  </sentences>
</xsl:template>
```

After this, I used a `for-each` statement that selects each `<sentence>` tag inside `<sentences>`:

```xml
<xsl:template match="/sentences">
  <sentences>
    <xsl:for-each select="sentence">
    </xsl:for-each>
  </sentences>
</xsl:template>
```

Looping over sentences, there's two values that are always present in every sentence which are `id` and `text`. These will be repeated regardless of whether we have an `<aspectTerms>` tag inside a given `<sentence>` tag. Therefore, it's reasonable to declare these as variables to be used later:

```xml
<xsl:variable name="id" select="@id"/>
<xsl:variable name="text" select="text"/>
```

These variables are declared inside the `<xsl:for-each select="sentence">` tag. First we take `id` from each `<sentence>`, which is done using `@id` in XPath. Then to take the text inside each `<text>` tag, I just reference the tag by name in XPath, since we're inside a given `<sentence>` tag, it can just be referenced using `text` in XPath.

Initially I had extracted this much less efficiently by traversing the entire XPath to each `text` tag matching the id variable like this: `<xsl:variable name="text" select="/sentences/sentence[@id=$id]/text/text()"/>`, the reasoning behind this was that I would match the contents of each `text` tag by id using `sentence[@id=$id]`, where `@id` is the id attribute of a sentence and `$id` is the reference to the id variable previously declared. Then the `text()` function outputs the contents of each `<text>` tag under the matched `sentence` tag.  This is NOT ideal, just use `text` (:

After, I identified 2 cases, those where I _did have_ an `<aspectTerms>` tag, and those where I didn't. Therefore, I used a `choose` statement. I could have used two `if` statements here, but the closest to an `if-else` here is a `choose-when-otherwise`, it makes the document a bit deeper than just using two `if` statements but IMO it's more readable like this. The first of the two statements is as follows:

```xml
<xsl:when test="aspectTerms">
  <xsl:for-each select="aspectTerms/aspectTerm">
    <sentence>
      <id><xsl:value-of select="$id"/></id>
      <text><xsl:value-of select="$text"/></text>
      <aspectTerm><xsl:value-of select="@term"/></aspectTerm>
      <from><xsl:value-of select="@from"/></from>
      <to><xsl:value-of select="@to"/></to>
      <polarity><xsl:value-of select="@polarity"/></polarity>
    </sentence>
  </xsl:for-each>
</xsl:when>
```

Here what I do is:

1. check using the `test` attribute of the `when` statement if the tag `<aspectTerms>` was present inside of the specific `<sentence>` tag the `for-each` loop referenced earlier.

2. If so, then I want to loop over each `<aspectTerm>` tag inside the contained `<aspectTerms>` tag that was found by the `when` statement.

3. Then, I create a `<sentence>` block with only top level tags inside of it, so that we can use `read_xml()` properly in pandas, since the resulting document needs to be flat. 

4. Inside each of the resulting `<sentence>` tags I add the `id`, `text` (variables I had declared previously, which in XSLT are referenced using `$`) and the `<aspectTerm>` properties `term`, `from`, `to` and `polarity`.

For the second case, the `otherwise` case (`else` in `if-else`) what I do is as follows:

```xml
<xsl:otherwise>
  <sentence>
    <id><xsl:value-of select="$id"/></id>
    <text><xsl:value-of select="$text"/></text>
  </sentence>
</xsl:otherwise>
```

This one is much simpler, as there's no `<aspectTerms>` tag, all I do is fill the `<sentence>` tag with the `id` and `text`, which I had previously declared as variables, hence using `$id` and `$text` to reference them.

And finally, the entire code to wrap it up:

```xml
<?xml version="1.0"?>
<xsl:stylesheet version="2.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
<xsl:template match="/sentences">
<sentences>
  <xsl:for-each select="sentence">
    <xsl:variable name="id" select="@id"/>
    <xsl:variable name="text" select="text"/>
    <xsl:choose>
      <xsl:when test="aspectTerms">
        <xsl:for-each select="aspectTerms/aspectTerm">
          <sentence>
            <id><xsl:value-of select="$id"/></id>
            <text><xsl:value-of select="$text"/></text>
            <aspectTerm><xsl:value-of select="@term"/></aspectTerm>
            <from><xsl:value-of select="@from"/></from>
            <to><xsl:value-of select="@to"/></to>
            <polarity><xsl:value-of select="@polarity"/></polarity>
          </sentence>
        </xsl:for-each>
      </xsl:when>
      <xsl:otherwise>
        <sentence>
          <id><xsl:value-of select="$id"/></id>
          <text><xsl:value-of select="$text"/></text>
        </sentence>
      </xsl:otherwise>
    </xsl:choose>
  </xsl:for-each>
</sentences>
</xsl:template>
</xsl:stylesheet>
```

This document, when combined with the XML, will generate the following transformation (utilizing the example shown at the top):

```xml
<?xml version="1.0"?>
<sentences>
  <sentence>
    <id>2339</id>
    <text>I charge it at night and skip taking the cord with me because of the good battery life.</text>
    <aspectTerm>cord</aspectTerm>
    <from>41</from>
    <to>45</to>
    <polarity>neutral</polarity>
  </sentence>
  <sentence>
    <id>2339</id>
    <text>I charge it at night and skip taking the cord with me because of the good battery life.</text>
    <aspectTerm>battery life</aspectTerm>
    <from>74</from>
    <to>86</to>
    <polarity>positive</polarity>
  </sentence>
  <sentence>
    <id>812</id>
    <text>I bought a HP Pavilion DV4-1222nr laptop and have had so many problems with the computer.</text>
  </sentence>
</sentences>
```


Which now that we successfully know works well after using the live server vscode extension to check by loading up the XML file in the web browser and checking the inspect element menu. With the XSL file in addition to the XML file inputted through the `read_xml()` function as follows:

```python
import pandas as pd

pd.read_xml('filename.xml', stylesheet='filename.xsl')
```

Results in the desired dataframe:

```
     id                                               text      aspectTerm  from    to  polarity
0  2339  I charge it at night and skip taking the cord ...            cord    41    45   neutral
1  2339  I charge it at night and skip taking the cord ...    battery life    74    86  positive
2   812  I bought a HP Pavilion DV4-1222nr laptop and h...            <NA>  <NA>  <NA>      <NA>
3  1316  The tech guy then said the service center does...  service center    27    41  negative
4  1316  The tech guy then said the service center does...    "sales" team   109   121  negative
```

And that's how you flatten XML files using XSLT!

***

## Additional resources

+ [W3school's XSLT tutorial](https://www.w3schools.com/xml/xsl_intro.asp)

+ [W3school's XPath tutorial](https://www.w3schools.com/xml/xpath_intro.asp)

+ [Excellent presentation by Rich from arbitrarytechnology on XPath and XSLT](https://youtu.be/WggVR4YI5oI)
