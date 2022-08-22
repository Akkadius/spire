## [1.13.5]

### Fixes

* [Task Editor] Fix issue with task activities saving properly

## [1.13.4]

### Misc Fixes

* [Task Editor] Fix an issue with `zone_version` field being populated with `null` on save of new activities
* [Merchant Editor] Fix an issue where the delete icon hover was showing "edit" instead of "delete"

## [1.13.3]

### New Task Fields Minor Update

* Add basic support for new task fields dz_template_id, lock_activity_id, replay_timer_group, request_timer_group, reward_point_type, reward_points
* More in depth selectors will be added for these fields later
* Fix the task editor's ability to save based on recent database schema changes

## [1.13.2]

### Alpha Minor Release

* Minor release that publishes early versions of NPC, Merchant editor etc.
* Support for legacy spell icons through user settings

## [1.13.1]

### UI Changes

Made changes to various components to streamline margins, cleaner UI across all parts of the app

Popups for Items and Spells now have more consistent behaviors, they stay within viewport, stay only on left or right of element - versus up/down/left/right, now have a zero delay, fade disabled to allow a cleaner experience 

Default windowing has been changed across the application as well.

Removed ghost margins in the scrollbar by overhauling most pages within the app to use proper element structure.

## [1.13.0]

### Basic Auth Support

If you want to run Spire on a hosted webserver with very basic authentication, you now can today until a more robust users, permissions system is built out.

Simply supply two environment variables `BASIC_AUTH_USER` and `BASIC_AUTH_PASSWORD` and Spire will only allow requests if you pass the basic authentication gate

## [1.12.0]

### Task Editor

A fully featured task editor is now available for use in BETA.

![image](https://user-images.githubusercontent.com/3319450/168457168-01227460-35fd-4a3b-a455-a17eca7db38b.png)

![image](https://user-images.githubusercontent.com/3319450/168457193-3be3ae97-e292-4f40-a248-cfbed36cadce.png)

![image](https://user-images.githubusercontent.com/3319450/168457197-2885598e-9c86-4a1b-966e-f66407c8a632.png)

![image](https://user-images.githubusercontent.com/3319450/168457202-0d777684-ad89-42f9-aaf3-01863125d220.png)

![image](https://user-images.githubusercontent.com/3319450/168457209-ebc86cf4-62fa-47a1-8866-52e3bbafeee4.png)

![image](https://user-images.githubusercontent.com/3319450/168457243-ae78e9e9-55a8-4216-8112-46326b794309.png)

![image](https://user-images.githubusercontent.com/3319450/168457258-5e3329d3-8dbd-4450-be89-d31950f7a821.png)

## [1.11.3]

### Relational Item Data

Item preview cards now display relational data, example below. As other tools are built out, this data can link out to those pages, for now it simply displays the data.

* Can be fished in
* Can be foraged in
* Dropped by
* Is a starting item for 
* Is found as a ground spawn
* Is sold by merchant(s)
* Is the result of tradeskill recipe(s)
* Is the reward of task(s)
* Unlocks doors

![image](https://user-images.githubusercontent.com/3319450/156943035-bb799182-9608-4e58-9a2a-88bf0a84947d.png)

### Misc

* Backend path management is now unified. Adds support for Spire being ran in a sub-folder ("bin" directory).

## [1.11.2]

### New Spell Icons

* Added over 2,000 spell icons from the live client, now available in the spell editor and used in any spell card
* Spell editor only shows up to 216 icons given that is what is actually used and available in RoF2

![image](https://user-images.githubusercontent.com/3319450/155861531-709650ba-a529-473e-ac6b-1ba311283d4e.png)

## [1.11.1]

### Item Card Previews

* Add model preview to top right of item card preview
* Edit button layout and UI changes

![image](https://user-images.githubusercontent.com/3319450/155836002-292931ba-b1d3-46fc-bf23-ac397a138c01.png)
![image](https://user-images.githubusercontent.com/3319450/155836156-3e228360-5713-410d-92be-6e707ec482fd.png)

## [1.11.0]

### Item, Spell, Database Strings Editor, Client Files and more

Item and Spell editor tools are now released in full, with the most feature rich sub-editors and selectors

### Item Editor

* Recast delay now displays in seconds instead of milliseconds
* Implemented sub-editors for the following fields: id, icon, idfile, material, color, augtype, proceffect, worneffect, focuseffect, scrolleffect, clickeffect, bardeffect
 
![Screenshot from 2022-02-26 01-06-08](https://user-images.githubusercontent.com/3319450/155833708-ab9a56d1-7e76-424a-ba23-1f0bdb46858e.png)
![Screenshot from 2022-02-26 01-07-23](https://user-images.githubusercontent.com/3319450/155833730-c39c8a6e-bf3a-4d21-887d-6420d4fb4f68.png)

### Spell Editor

* Added data translations to most remaining fields to make them more obvious to users
* Added formulas to SPA effects
* Conditionally display fields that are not contextually relevant unless certain fields are set
* Fix buff duration display not properly showing times
* Implemented a cone visualizer to real-time preview cone spells
* Implemented a horse selector for **teleport_zone** via horse type spells
* Implemented a pet selector for **teleport_zone** via pet type spells
* Implemented a spell selector for spell related fields
* Implemented a zone selector for **teleport_zone** via pet translocate / gate spells
* Implemented an item selector for item related fields
* Implemented buff duration formulas
* Implemented smart Spell and Item selection inside the effects tab, where SPA fields base, limit, max will pulsate green depending on the SPA to indicate item or spell selection
* Implemented sub-editors for the following fields: id, casting_anim, target_anim, icon, spellanim, cone_start_angle, cone_stop_angle, nimbuseffect, aoerange, teleport_zone, range, min_range, typedescnum, effectdescnum, effectdescnum_2, descnum, components_1, components_2, components_3, components_4, noexpend_reagent_1, noexpend_reagent_2, noexpend_reagent_3, noexpend_reagent_4, recourse_link
* SPA effects are now highlighted when hovered over and selected
* When an SPA is selected, the editor will select reference values from another spell in the database with the same SPA

![Screenshot from 2022-02-26 01-05-12](https://user-images.githubusercontent.com/3319450/155833675-ecf575a5-2dd1-404e-a1c2-0988a65504ca.png)
![Screenshot from 2022-02-26 01-03-54](https://user-images.githubusercontent.com/3319450/155833645-cbd9c839-6d43-4ee0-b505-9425381eccbb.png)

### Item / Spell Editors

* Added cast bar simulation components to any field that references cast times
* Fixed sub-editor highlight selection scroll consistency issues where it would not always scroll to the selected preview
* Implement database column filtering that is additive to the normal search fields
* Now have cloning functionality
* Search improvements, adding table and card toggles, result limit toggles
* Sub-editor fields now pulsate green to indicate to the user that there is a sub-editor associated with the field
* Sub-selectors now only activate during input click instead of on hover
* Video viewers now display a black background by default before the video loads in

### Database Strings Editor

* Implemented a database strings editor that will allow users to list, create, edit and delete database strings per string type

![Screenshot from 2022-02-25 23-19-53](https://user-images.githubusercontent.com/3319450/155830410-757d1a27-d7ad-419a-9805-a197fbffae3d.png)

### Client Files Importer / Exporter

![image](https://user-images.githubusercontent.com/3319450/155830515-e74d1b27-f2a0-4fa6-afcc-c6c93e70cd8a.png)

* Implemented client files management supporting both spells_us.txt and db_str.txt

### General

* Added a range visualizer component that will visualize in game ranges of up to 1,000 units
* All usage of Spell icons in the application have been replaced with sprites
* Item card preview now shows factions

## [1.6.7]

### Quest API Explorer

Implemented searching by method param. [#23](https://github.com/Akkadius/spire/issues/23)

![image](https://user-images.githubusercontent.com/3319450/151752471-85e4df08-4f0c-4936-a082-f62557d1cd12.png)

## [1.6.6]

### UI Tweaks

All item icons in application now load from a spritesheet, removing the dependency on an additional set of 40MB of images being loaded into the binary

## [1.6.5]

### Quest API Explorer

* Fixed a bug where sub event examples for Lua were not formatted properly [#21](https://github.com/Akkadius/spire/issues/21)

### Viewer Improvements

* Viewers now show separate loaders for when content is being downloaded versus when content is being rendered

![image](https://user-images.githubusercontent.com/3319450/151697911-f25ab05c-1229-45c5-9a30-9069d32eb318.png)

![image](https://user-images.githubusercontent.com/3319450/151697961-de014617-4cdf-4c0f-99dc-1c01f7b1690c.png)

* Viewers now don't show content until images are fully downloaded and rendered
* Item Icon Viewer now has icons up through live

### Spire Binary Size

* Spire should now be roughly 100MB less in executable size.

## [1.6.4]

### Viewer Improvements

* Icon viewer inputs are now separated from the preview pane
* Item model viewer inputs are now separated from the preview pane
* Icon viewer now has a fixed scroll viewport
* Item model viewer now has a fixed scroll viewport
* Emitter viewer now has a fixed scroll viewport
* Spell animations viewer now has a fixed scroll viewport
* Video viewers are by default larger to improve browser rendering performance and make it easier to digest visual scanning

## [1.6.3]

### Viewer UI Tweaks

* Various UI tweaks to viewers for a much cleaner, simple user experience

## [1.6.2]

### Spell Animation Previewer

* Fixed bug that prevented all previews from showing properly

## [1.6.1]

### Linux Auto Updates

* Linux auto updates should properly work from this release (1.6.1) and on. 

## [1.6.0]

### Updated Spell Animations!

* Special thanks to DeadZergling for all of his effort putting together these high quality preview videos

[![](https://img.youtube.com/vi/7xXoBAWs3n8/0.jpg)](https://www.youtube.com/watch?v=7xXoBAWs3n8)

* Race viewer initial load should now be faster
* Race viewer now includes a models available by zone filter

## [1.5.1]

* Changelog page now only autoplays videos that are in the current viewport
* Adjusted preview video size for emitters
* Improvements to video preview rendering logic

## [1.5.0]

### Player Animation Viewer

* Spire now has an player animation viewer that can be used standalone and in things like a Spell editor where there are
  casting and target animations when spells are casted
* Special thanks to DeadZergling for all of his effort putting together these high quality preview videos

[![](https://img.youtube.com/vi/_WLjso1d9p8/0.jpg)](https://www.youtube.com/watch?v=_WLjso1d9p8)

## [1.4.0]

### Emitter Viewer

* Spire now has an emitter viewer that now makes it easy for anyone looking to manipulate emitter text files in the
  client to find what kind of effects they want to play with or insert into their zones
* Special thanks to DeadZergling for all of his effort putting together these high quality preview videos

[![](https://img.youtube.com/vi/dpE_xWR2i6o/0.jpg)](https://www.youtube.com/watch?v=dpE_xWR2i6o)

## [1.3.1]

### Spire Self Updating

* This release implements self updating and notifies users when a new version of Spire is available for download

## [1.3.0]

### Spire Launch!

**Akkadius**

Spire is now available for general release and the repository can be found
here [https://github.com/Akkadius/spire](https://github.com/Akkadius/spire)

### Spire Desktop Release

Spire can be ran as a standalone executable on Windows.

Download the latest release
here [https://github.com/Akkadius/spire/releases](https://github.com/Akkadius/spire/releases)

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

* Item editor comes bundled with many hover selector and preview tools to make development feedback loop the tightest
  and best experience it can possibly be. Examples of these include:
* Item free ID selector (hover over "id")
* Item model selector (hover over "idfile")
* Item icon selector (hover over "icon")
* Item color selector (hover over "color")
* Item material previewer (hover over "material", arrow selector) (cloth, chain, plate, robe etc.)
* Item augment type selector (hover over "augtype")
* Item effects selector (hover over any effect "spell id" field)

#### Item Editor Sane Defaults

* The item editor does a lot in the way of setting sane defaults for the user. There are a lot of fields that need to be
  set when other fields are set so there are triggers that are setup to make this more intuitive.
* When **worneffect** is set to a non-zero value, **worntype** is set to 2
* When **clickeffect** is set to a non-zero value, **clicktype** is set to 5, recast delay set to 12000, maxcharges -1
  and casttime to 3000
* When **scrolleffect** is set to a non-zero value, **scrolltype** is automatically set to 7
* When **bardeffect** is set to a non-zero value, **bardeffecttype** is set to 8
* When **augslot_1_type** is set to a non-zero value, **augslot_1_visible** is set to 1
* When **augslot_2_type** is set to a non-zero value, **augslot_2_visible** is set to 1
* When **augslot_3_type** is set to a non-zero value, **augslot_3_visible** is set to 1
* When **augslot_4_type** is set to a non-zero value, **augslot_4_visible** is set to 1
* When **augslot_5_type** is set to a non-zero value, **augslot_5_visible** is set to 1

### Quest API Explorer

* Added a **Refresh** button for local and desktop builds for anyone who wants to refresh their Quest definitions
  manually. Production and hosted version of Spire receives webhooks when new commits are made to EQEmu/Server and local
  installations won't receive this and require a manual update.

### General Fixes

* (Kinglykrab reported) Dark Elf short name is now DEF, Wood Elf short name is now ELF, High Elf short name is now HIE,
  Halfling short name is now HFL
* (Kinglykrab reported) Powersource icon was using Charm icon and is now the proper icon
* Scrollbar should be twice the size in width now
* (Trust reported) Expansion bitmask calculator now reports 0 (Classic) correctly
* Added an **Only** button to classes in the items search view so that search results can filter on items that are only
  equipped by selected class.

## [1.0.3]

Add version to navbar

## [1.0.2]

Embed the changelog you're reading into the app

## [1.0.1]

Remove application boot dependency on having `APP_NAME` set

## [1.0.0]

Initial release
