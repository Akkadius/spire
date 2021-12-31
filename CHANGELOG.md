## [1.3.0]

### Spire Launch!

**Akkadius** 

Spire is now available for general release and the repository can be found here [https://github.com/Akkadius/spire](https://github.com/Akkadius/spire)

### Spire Desktop Release

Spire can be ran as a standalone executable on Windows. 

Download the latest release here [https://github.com/Akkadius/spire/releases](https://github.com/Akkadius/spire/releases)

To run, simply put **spire.exe** in your server directory and double click (On Windows)

### Item Editor v1 Launch!

[![](https://img.youtube.com/vi/iQvm9pH8420/0.jpg)](https://www.youtube.com/watch?v=iQvm9pH8420)

* Initial set of item search filters, more to be added later down the road
* Search for items by either **table** or **card** formats, **table** is default
* Item card preview on table item row hover
* All item fields are supported out of the box with a detailed editor
* Class bitmask selector
* Race bitmask selector
* Deity bitmask selector
* Slot bitmask selector

#### Hover Selector / Preview Tools

* Item editor comes bundled with many hover selector and preview tools to make development feedback loop the tightest and best experience it can possibly be. Examples of these include:
* Item free ID selector (hover over "id")
* Item model selector (hover over "idfile")
* Item icon selector (hover over "icon")
* Item color selector (hover over "color")
* Item material previewer (hover over "material", arrow selector) (cloth, chain, plate, robe etc.)
* Item augment type selector (hover over "augtype")
* Item effects selector (hover over any effect "spell id" field)

#### Item Editor Sane Defaults

* The item editor does a lot in the way of setting sane defaults for the user. There are a lot of fields that need to be set when other fields are set so there are triggers that are setup to make this more intuitive.
* When **worneffect** is set to a non-zero value, **worntype** is set to 2
* When **clickeffect** is set to a non-zero value, **clicktype** is set to 5, recast delay set to 12000, maxcharges -1 and casttime to 3000
* When **scrolleffect** is set to a non-zero value, **scrolltype** is automatically set to 7
* When **bardeffect** is set to a non-zero value, **bardeffecttype** is set to 8
* When **augslot_1_type** is set to a non-zero value, **augslot_1_visible** is set to 1
* When **augslot_2_type** is set to a non-zero value, **augslot_2_visible** is set to 1
* When **augslot_3_type** is set to a non-zero value, **augslot_3_visible** is set to 1
* When **augslot_4_type** is set to a non-zero value, **augslot_4_visible** is set to 1
* When **augslot_5_type** is set to a non-zero value, **augslot_5_visible** is set to 1

### Quest API Explorer

* Added a **Refresh** button for local and desktop builds for anyone who wants to refresh their Quest definitions manually. Production and hosted version of Spire receives webhooks when new commits are made to EQEmu/Server and local installations won't receive this and require a manual update.

### General Fixes 

* (Kinglykrab reported) Dark Elf short name is now DEF, Wood Elf short name is now ELF, High Elf short name is now HIE, Halfling short name is now HFL
* (Kinglykrab reported) Powersource icon was using Charm icon and is now the proper icon
* Scrollbar should be twice the size in width now
* (Trust reported) Expansion bitmask calculator now reports 0 (Classic) correctly
* Added an **Only** button to classes in the items search view so that search results can filter on items that are only equipped by selected class.

## [1.0.3]

Add version to navbar

## [1.0.2]

Embed the changelog you're reading into the app

## [1.0.1]

Remove application boot dependency on having `APP_NAME` set
 
## [1.0.0]

Initial release
