## [2.5.4]

### Lua Object Return Methods

- Fixes Lua Object Return Methods not showing in Lua Methods

## [2.5.3]

### Spire Application Icon

- Application icon has been added to windows

## [2.5.2]

### Downloader Optimizations

- Made some optimizations to the downloading of assets and bootstrapping of assets

## [2.5.1]

### Windows Fixes

- Resolved issues with windows downloading logic
- Resolved issues with windows path resolution in static asset serving

## [2.5.0]

### Spire Binary Size Reduction

- Spire is now 26MB compressed and 55MB raw from original 516MB compressed and 738MB raw. 13x reduction in binary size
- Static assets are now decoupled from the executable binary and are now downloaded once during the first initial boot of Spire. They won't ever need to download again unless assets update.
- Static assets will also automatically update when there are new assets.
- This will make Spire releases faster, the distributed binary leaner and also result in less wasted bandwidth

## [2.4.2]

### [Items] Ammo / Powersource

- Fixed a minor issue where ammo / power source slots were flip-flopped in the item editor.

## [2.4.1]

### [Bots] Minor fix (Aeadoin)

- Fixed an issue when adding new Spells to the Bot Spell List.

## [2.4.0]

### [Bots] Spell Editor (Aeodoin)

- Added Bot Spells Entries database relationships to NPC Spells, and Spells New
- Added Bot Nav Tab
- Added Bot Spells under Bot Nav
- Bot Spells pulls Bot Spell Data from bot_spells_entries table
- Same features/functionality as NPC Spells (View spell sets, edit, add new spells to spell set, delete spells from spell set) I removed the ability to delete the Bot Spell list.

![image](https://user-images.githubusercontent.com/3319450/201459665-f92080ee-e730-4737-a610-a921de89cbb5.png)

![image](https://user-images.githubusercontent.com/3319450/201459702-962c8dd9-1d28-4492-b007-bb5ca7a1de38.png)

## [2.3.4]

### [Tasks] Fix Cloning with Activities

Fixed an issue where task cloning did not fully clone all associated activities

## [2.3.3]

### [Items] Negative Values Preview

Item preview renders now properly show negative values.

![image](https://user-images.githubusercontent.com/3319450/197303792-2c4a4404-f53a-4955-bf73-1850cbf117a2.png)

![image](https://user-images.githubusercontent.com/3319450/197303805-4d089d23-906a-441f-a77b-3e2864a7bf72.png)

## [2.3.2]

### [Items] Click Type Naming

Click Types were named improperly and not fully descriptive.

## [2.3.1] 

### [Cloning] Minor fix

Fix issue where older database versions do not support the MySQL clause `RETURNING id`. This disables the newer syntax allowing creates to work properly again.

## [2.3.0]

### [Database Connections] Users, Permissions & Audit Logging

### Users  

Users can now be added to server connections that you own. Any user that has a user on Spire, hosted or local is available to add to a connection.

![image](https://user-images.githubusercontent.com/3319450/196054882-99f6f778-8844-4362-8c4c-708bafc715e4.png)

![image](https://user-images.githubusercontent.com/3319450/196054944-40816080-fe7e-4e25-8692-a9d63db3713a.png)

### Permissions

With users, come permissions. You are now able to manage permissions with both read and write for every resource available through Spire. You can also use **ALL** to grant users to all read or write.

![image](https://user-images.githubusercontent.com/3319450/196055017-925c6188-d7fa-4cf9-8af1-17f66a93281e.png)

### Audit Logging

With users (Developers) come the need to monitor bad actors and mistakes!

![image](https://user-images.githubusercontent.com/3319450/196055080-67e25106-9fb6-45aa-83d5-f92e95a3bc97.png)

View audit logs

![image](https://user-images.githubusercontent.com/3319450/196055087-9cf119b3-c496-470a-8a1d-fae4813395f8.png)

#### Discord Webhook Audit Logging

![image](https://user-images.githubusercontent.com/3319450/196055054-0578b668-fd4a-43e2-8dce-4762b29cfdd1.png)

![image](https://user-images.githubusercontent.com/3319450/196055117-a70d2d7b-21c7-4d32-ab75-f369653843e0.png)

### Misc Changes

* Item editor now has a delete button
* Spell editor now has a delete button
* Fix expansion offsets to no longer be offset by 1 per recent server changes
* Fix cash display to be more properly horizontally aligned

## [2.2.5]

### [Item Editor] Item Caching

* Fixed issue where item being cloned and reassigned to new ID's would reflect wrongly display after save

## [2.2.4]

### [Task Editor] Update Logic Fixes

* Fix issue with nested objects showing up in save actions, preventing some things from saving properly

## [2.2.3]

### [Task Editor] Save Action Freezing

* Added save action freezing to inputs and buttons when work is being done to change the database. This will ensure that actions are complete before subsequent actions can potentially mess up task data.
* Added a "Fix and re-order task activities" button to fix broken task activities ordering if it were to be in a bad state. Button does nothing if the task is correct

## [2.2.2]

### [Update Logic] Handle Nullables

Resolved an issue where nullable fields were not properly being picked up in new update logic.

## [2.2.1]

### [Update Logic] Only Update What's Changed

Update logic for all saves has been rewritten to only update fields that have changed. This will also as a side effect resolve an issue where when values are set to zero, the object would not properly save.

## [2.1.1]

### [Items] Item Slot Ordering

Item slots "Ammo" and "Power Source" were swapped.

## [2.1.0]

### [Task Editor] Task Preview Requisite Activity ID

Adds support for task requisite activity ID.

Tasks for example in Dragons of Norrath have activities that are not sequential and can be unlocked early in task progression. This update adds support on the editor side.

![image](https://user-images.githubusercontent.com/3319450/193499479-657c0a76-0b49-4509-a324-10524f3b11f9.png)

## [2.0.3]

### [Task Editor] Minor Fixes

* Fix an issue where `item_id_list` was binding to an integer instead of a string
* Adjust UI elements in sub-editor panes (bordered tables)

## [2.0.2]

### [Connection Manager]

Fix an issue where you weren't able to switch active connections

## [2.0.1]

### [Merchant Editor] Minor fix

Fix issue where Merchant Editor wasn't rendering NPC preview cards properly

## [2.0.0] 

### [Task Editor] Major Task Update(s)

![image](https://user-images.githubusercontent.com/3319450/192701596-13a3fbbe-7c46-47ca-8f54-16de57c3c8d9.png)

### Changes

* Added sectional descriptions and info indicators
* Description building should be more clear and concise
* Explore details are now inline with the activity
* Fields `item_id` and `item_goal_id` combine into `item_id_list` which can contain a single item or list of items that are pipe delimited `|` Example (37025|37029|37032)
* Fields `npc_id` and `npc_goal_id` combine into `npc_match_list` which can contain a single NPC or a list of NPC's that are pipe delimited `|` Example (4007|4009|4013|4024|4036|4043|4078|4080) and can also contain NPC names as partials (orc|gnoll|bear)
* Instead of `goalmethod` and `rewardmethod` displayed as dropdowns. It simply is a checkbox that displays **Quest Controlled**
* Performance and caching updates
* Resolved issues with tasks saving inconsistently
* Reward window should now display multiple items
* Task rewards now support multiple items inline
* Tasks now save automatically when a field is modified and activity actions are invoked
* Tasks now set proper defaults for `req_activity_id` (-1)

See Server PR for more major system change details https://github.com/EQEmu/Server/pull/2449

### NPC Match List Support

Match lists are **| pipe separated lists** containing NPC ID(s) or partial (wildcard) NPC names. This is very useful especially when trying to filter activity updates by name.

NPC match lists are available with the following activity types

* Deliver
* Kill
* Speak with
* Give

Fields `npc_id` and `npc_goal_id` have combined 

![image](https://user-images.githubusercontent.com/3319450/192702240-025ba8ac-c09f-458e-ae29-73d8d3379e03.png)

### Item Match List Support

To bring similar flexibility to that there of **NPC Match List(s)**, multiple item's can be used for the following activity types.

* Deliver
* Loot
* Tradeskill
* Fish
* Forage

![image](https://user-images.githubusercontent.com/3319450/192705534-d9f02103-1b2e-4f53-8eb9-c2cc46ce7e00.png)

### Exploration Box Changes

Exploration used to be managed by the **proximities** table. Part of the recent waves of simplification this is now managed on the **task_activities** table to simplify. As a result explore boxes are now visible within the activity pane instead of breaking out as their own sub-editor when editing an explore activity type.

![image](https://user-images.githubusercontent.com/3319450/192704930-aa0efda8-b196-43a5-8311-755481327763.png)

### Multiple Item Reward(s)

Not only are multiple rewards now possible through the **Reward Item ID(s)** field (**separated by |**) you can also see them render in the **Task Preview Window**

![image](https://user-images.githubusercontent.com/3319450/192703822-f234928b-4e89-4f65-b7de-9123cd90ff48.png)

### Field Usage Matrix

To be updated in official documentation

| activity\_type | activity\_type\_description | item\_id\_list | npc\_match\_list | dz\_switch\_id                    | min\_x | max\_x | min\_y | max\_y | min\_z | max\_z | skill\_list   | spell\_list   | item\_list                               | target\_name                           | Description Format                                 |
| -------------- | --------------------------- | -------------- | ---------------- | --------------------------------- | ------ | ------ | ------ | ------ | ------ | ------ | ------------- | ------------- | ---------------------------------------- | -------------------------------------- | -------------------------------------------------- |
| 1              | Deliver                     | x              | x                | \-                                | o      | o      | o      | o      | o      | o      | \-            | \-            | Name of Item(s)                          | Name of NPC                            | Deliver (goalcount) (item\_list) to (target\_name) |
| 2              | Kill                        | \-             | x                | \-                                | o      | o      | o      | o      | o      | o      | \-            | \-            | \-                                       | Name of NPC                            | Kill (goalcount) (target\_name)                    |
| 3              | Loot                        | x              | o                | \-                                | o      | o      | o      | o      | o      | o      | \-            | \-            | Name of Item(s)                          | Name of NPC                            | Loot (goalcount) (item\_list) from (target\_name)  |
| 4              | Speak With                  | \-             | x                | \-                                | o      | o      | o      | o      | o      | o      | \-            | \-            | \-                                       | Name of NPC                            | Speak with (target\_name)                          |
| 5              | Explore                     | \-             | \-               | \-                                | x      | x      | x      | x      | x      | x      | \-            | \-            | \-                                       | Name of Explore Area                   | Explore (target\_name)                             |
| 6              | Tradeskill                  | x              | \-               | \-                                | o      | o      | o      | o      | o      | o      | \-            | \-            | Name of Item(s)                          | \-                                     | Create (goalcount) (item\_list) using tradeskills  |
| 7              | Fish                        | x              | \-               | \-                                | o      | o      | o      | o      | o      | o      | \-            | \-            | Name of Item(s)                          | \-                                     | Fish for (goalcount) (item\_list)                  |
| 8              | Forage                      | x              | \-               | \-                                | o      | o      | o      | o      | o      | o      | \-            | \-            | Name of Item(s)<br>Forage x (item\_list) | \-                                     | Forage (goalcount) (item\_list)                    |
| 9              | Use (Cast On)               | \-             | \-               | \-                                | o      | o      | o      | o      | o      | o      | \-            | Spell name(s) | \-                                       | Name (Anything)                        | Use (spell\_list) on (target\_name)                |
| 10             | Use2 (Skill On)             | \-             | \-               | \-                                | o      | o      | o      | o      | o      | o      | Skill name(s) | \-            | \-                                       | Name (Anything)                        | Use (skill\_list) on (target\_name)                |
| 11             | Touch                       | \-             | \-               | x<br>(doors table dz\_switch\_id) | o      | o      | o      | o      | o      | o      | \-            | \-            | \-                                       | Name of Touch Target<br>(Touch target) | Touch (target\_name)                               |
| 100            | Give                        | \-             | x                | \-                                | o      | o      | o      | o      | o      | o      | \-            | \-            | \-                                       | Name of NPC                            | Give (goal\_count) to (target\_name)               |
| 255            | Quest Script                | \-             | \-               | \-                                | o      | o      | o      | o      | o      | o      | \-            | \-            | \-                                       | \-                                     | \-                                                 |

## [1.13.10]

* Addresses an issue where connections endpoint was hit when booting the app. When Spire does not have a Spire database, it panics the backend. https://github.com/Akkadius/spire/issues/63

## [1.13.9]

### [Task Editor] Minor Update

* Minimal update to restore task editor functionality from schema changes https://github.com/EQEmu/Server/pull/2402
* Proximity functionality will be broken until new fields are handled

## [1.13.7]

### [NPC Spells Editor] Early alpha release

* Automatic spell type detection
* List previews
* Manage NPC spell sets fully
* NPC preview
* Search
* Spell list nesting highlighting
* Table pagination

## [1.13.6]

### Fixes

* [Merchant Editor] Fix issue with new items being added not having proper defaults keeping Merchant items from loading properly

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
