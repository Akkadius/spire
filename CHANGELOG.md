## [3.13.1] 2/11/2024

**Sage** Introduce new page for integration with EQ Sage. Sage, among other things, will be a 3D zone database editor that plugs into the existing Spire API. It is currently in a prototyping/experimental phase and will be undergoing changes to initially support zone editing in a broad sense with details for each individual aspect explored and fine tuned to provide full support for an immersive editing experience. More info can be found in the EQEmu Discord under the #project-requiem channel and any questions can be directed to the creator and maintainer, temp0.

#### Spawn Entry Editing, Pathfinder

![image](https://github.com/Akkadius/spire/assets/3319450/29f34cb6-6961-4287-aadd-cd99a3644c5b)

#### S3D / EQG Zone Region Inspector

![image](https://github.com/Akkadius/spire/assets/3319450/1c88abc3-25a9-4562-a869-976fee530a44)

## [3.12.3] 2/7/2024

* **Spire Admin** Fix issue where the server update page would not load because it was not rejecting errors from https://spire.akkadius.com properly. The error occurred because the hosted site is now **https** (SSL) and we hadn't updated our fetch logic to handle this. This is now fixed and the server update page will load properly again.

## [3.12.2] 2/4/2024

* **Database** Implement log database connection resolving. This allows operators to utilize the "QS" database connection type for logs, particularly sinking player event logs to an external database or server. When a QS connection is specified in the `eqemu_config.json` - player event logs will automatically route to that connection. This is useful for operators who want to sink logs to an external database or server for auditing purposes. The database configuration gets synchronized and injected during Spire bootup and also during Spire Admin configuration saving.
* **Player Event Log Viewer** Player event log request is split from character data bulk loading two split the queries on separate connections
* **Spire Admin** When making configuration changes in the configuration editor UI - the database connection resolver will flush along with any cached user connections. The `eqemu_config.json` will also be re-injected into the `spire_server_database_connections` table
* **Database Resolver** Added a connection type of `eqemu_logs` of which `player_event_logs` and `player_event_log_settings` both route to currently.
* **Configuration Sync** Ensure we also empty fields when they are removed from the `eqemu_config.json` when saving the configuration in Spire Admin

## [3.11.3] 1/29/2024

* **Spire Admin** Message of the day should now be functional again
* **Spire Admin** Spire now has a server lock / unlock button and status icon that is now present in the admin header visible on all admin pages.
* **Spire Admin** The dashboard page now has a "Server Addresses" section that displays server world address and local address with copy to clipboard functionality if they are present in the server configuration.
* **Spire Admin** The dashboard page now has a "Networking Interfaces" section that displays interfaces and realtime bandwidth utilization.
* **Special Abilities Calculator** Adds support in Special Abilities Calculator for `Immune to Assassinate` and `Immune to Headshot` abilities. (Kinglykrab)
* **Spire** Add in application update notifications that inform the user when there is an update available. This is a new feature that will be used in the future to notify users of new Spire releases while in the UI. It supports ignoring the update, only checks periodically.
* **Spire** Add a manual update check button to the navbar.

![image](https://github.com/Akkadius/spire/assets/3319450/b0576ea5-0ce4-4062-a456-768985faca4b)

## [3.9.11] 1/24/2024

* **Installer** Add back auto installing VC runtime 2015-2019 for Windows users still using Windows 10
* **Installer** Shorten the server name randomized string to 5 from 10 to prevent UI display issues and make it easier to identify

## [3.9.9] 1/23/2024

* **Task Editor** Fix issue where new tasks were not getting inserted with enabled `1`

## [3.9.8] 1/15/2024

* **Spire Admin** Fix issue where editing server configuration would create empty loginserver entries
* **Spire Admin** Added more resiliency around server configuration editing for loginserver blocks. If loginserver 3 is defined and loginserver 2 is not, it will now create loginserver 2 and fill it with default values. Same goes for position 2 and 1. This will prevent the server from failing to boot due to missing loginserver blocks.

## [3.9.7] 1/13/2024

* **Database Connections** Fix core issue database connection management logic where database connections that failed to established could be cached and used in subsequent requests
* **Database** Update Spire with latest EQEmulator schema changes
* **Spire Admin** Tweaks to the UI to keep parts of the UI jumping around while the page is loading
* **News** News page now takes the full page on initial load
* **Installer** Randomize long name for server installs to prevent collisions

## [3.9.1] 1/7/2024

* **Spire Install** Fix issue during Spire initialization where users were not able to disable authentication receiving an error

## [3.9.0] 1/6/2024

* **Server Config** Update server admin configuration editor to use new config schema. Mailserver and chatserver have been consolidated to `server.ucs` see https://github.com/EQEmu/Server/pull/3768

## [3.8.8] 1/6/2024

* **Server Installer** Use legacy analytics pipeline to record install counts

## [3.8.7] 12/31/2023

* **Server Installer** Run Spire as regular user on Windows instead of admin
* **Server Installer** Sign windows binaries

## [3.8.4] 12/30/2023

* **Downloader** Tweaks to downloader logic

## [3.8.3] 12/29/2023

* **Server Installer** Add Windows run as admin check to installer
* **Server Installer** Add Windows firewall rule creation to installer
* **Server Installer** Add Windows cmd prompt window maximization to installer
* **Server Installer** Add "Press enter to continue" to installer when an error occurs to keep the window from disappearing before the user can read the error
* **Server Installer** Add "Press enter to continue" to installer at the end of installation
* **Database** Set connection max lifetime to 3 minutes to prevent stale connections from being used

## [3.8.1] 12/29/2023

* **Downloader** Add more resiliency to the file downloader so it fully flushes the file before closing the file handle.

## [3.8.0] 12/29/2023

* **Login Automation** Add ability to login to the login page by passing query params `user` `password` and `redirect`

## [3.7.10] 12/27/2023

* **Spire Launcher Proxy** Speed up command line launcher proxy calls to start/stop/restart

## [3.7.9] 12/27/2023

* **Spire Updates** Fix an issue where Spire command invocations were making API requests to GitHub to check for new versions when they didn't need to.
* **Utility** Added a `HTTP_DEBUG` environment variable flag that will print exhaustive request debugging
* **Installer** Start to build eqemu server installer v2 binaries alongside of releases. These are not ready for prime time yet but they will be!

## [3.7.8] 12/12/2023

* **Items** Fix an issue with the items page loading due to a recent starting items schema change.
* **Items** Items search page now surfaces error to the user.

## [3.7.7] 11/21/2023

* **Server Updates** Adjust visual feedback when updating server binaries via CLI

## [3.7.6] 10/31/2023

* **Spire Updater** Fix issue with Spire updater not prompting for any key to restart when there is no terminal attached

## [3.7.5] 10/31/2023

* **Player Event Logs** Fix issue with player event log loading in some cases.

## [3.7.3] 10/29/2023

* **DiscordEQ** Make the DiscordEQ section of **eqemu_config.json** optional

## [3.7.2] 10/22/2023

* **DiscordEQ** Add **eqemu_config.json** configuration support for legacy DiscordEQ 

## [3.7.1] 10/19/2023

* Installer release binary test

## [3.7.0] 10/18/2023

**EQTraders** Added eq traders CLI tools to both scrape and import
 * `eq-traders:scrape` will scrape eqtraders.com and output a JSON file. It will cache all item lookups and page lookups and store them under `./data/eqtraders`
 * `eq-traders:import` will import the JSON file into the database

## [3.6.0] 10/18/2023

* **Assets** Keep from checking GitHub for asset updates every boot. We check once an hour now on boot and during regular application updates. This prevents us from getting rate limited by GitHub which can prevent other Spire actions from working properly.
* **Launcher** `spire:spire:launcher stop|restart` will now also kill parent wrapper processes

## [3.5.7] 10/16/2023

* **Zip** Ignore zip deletion errors when deleting old zips
* **Spire Admin** Fix race condition when rendering server header card
* **Spire Admin** Fix race condition during log file streaming when switching between files
* **Spire Admin** Add CLI command `eqemu-server:update` to handle updating an eqemu server via release or self-compiled 

## [3.4.1] 10/16/2023

* **Crash Analytics** Fix fingerprinting false positives in windows crash analytics
* **Crash Analytics** Add loader to page
* **Auth** Log users out in cases where their token is invalid. Also catches cases where encryption key was changed.

## [3.4.0] 10/15/2023

* **Crash Analytics** Added fingerprinting to crash analytics to help identify duplicate crashes.
* **Crash Analytics** Add Discord webhook to crash analytics to notify developers of new unique crashes.
* **Server Updater** Added unique crash counts to the server updater page within Spire Admin

![image](https://github.com/EQEmu/Server/assets/3319450/3d7dd8ec-2269-4041-a22a-afb57395a56c)

## [3.3.5] 9/2/2023

* **Video Renderer** Additional performance, responsiveness, memory, and stability improvements to the video renderer.

## [3.3.4] 9/2/2023

* **Spell Editor** ID selector input now only engages on input focus instead of hover.
* **Video Renderer** Optimizations made to the video renderer to fully destroy video elements when out of view, not just stop them from playing. This will help with memory usage and performance. This will benefit all video render standalone utilities and sub-editor panes.
* **Database Backup (Spire Admin)** Fix an issue with database backup form where compress would reset other selections.
* **Spell Animations** Default spell previews to start at 3 seconds of the video

## [3.3.3] 9/2/2023

Remove Alpha setting toggle and leave alpha tools on by default but still show their badge in navigation as alpha.

## [3.3.2] 8/19/2023

Replace Quest API explorer VSCode extension embed with a static image that links to the extension in marketplace

## [3.3.1] 7/31/2023

Add zone log streaming back into admin panel zone servers page

## [3.3.0] 7/19/2023

Fix a critical issue where saving new entities (items or spells for example) would also duplicate all relational data 

## [3.2.2] 7/6/2023

### Tasks

Fix an issue where the task list wasn't being refreshed when a task was saved

### General

Fix an issue with the free id fetcher logic when a table uses completely contiguous ranges, we will default to the max id if there are no ranges available

## [3.2.1] 7/5/2023

* Spire local setups will now update the database connection name in the connections page on bootup when the server long name has changed
* Spire local setups will now update the database connection name in the connections page when the server long name has changed and the server is running

## [3.2.0] 7/4/2023

* When updating server binaries, inform users that they need to power off their server before attempting to update binaries
* Extend `spire:init` command to take two new flags `--compile-build-location=/home/eqemu/code/build/` and `--compile-server=true|false`. This is used during the eqemu server installer to determine whether Linux users want to initialize their server configuration with `release` or `self-compiled`

## [3.1.10] 7/4/2023

* Fix display issue with Spire Admin header when sized to different sizes
* Fix edge case issue on Windows with telnet connections to world causing state issues in the dashboard, such as not showing zone list

## [3.1.9] 7/3/2023

* Display warning to windows users during server start about windows spawning

## [3.1.8] 7/3/2023

* Fix issue where `minZoneProcesses` is being sent as a string instead of a number in the launcher window

## [3.1.7] 7/3/2023

* Fix issue where the login screen would be scrunched up on mobile devices
* Fix mobile issues in the server admin UI in the header
* Fix server update releases page to be responsive with header content windows
* Fix an issue where the navbar collapse would not display correctly on mobile
* Fix an issue where Spire release page notifications would appear on the bottom of the page versus the top
* Fix an issue where the Spire release page would not function properly if `spire.akkadius.com` was offline (defaults to local if hosted is not online)

Ignore frequently called admin endpoints in the Spire console output to make it easier to see relevant endpoints being called

```
/api/v1/admin/system/cpu
/api/v1/admin/system/resource-usage-summary
/api/v1/eqemuserver/client-list
/api/v1/eqemuserver/server-stats
```

## [3.1.6] 7/3/2023

* Removed "Failed to connect to gameserver" error when navigating to server admin page "Zone Servers" and the server is offline
* Removed "Failed to connect to gameserver" error when navigating to server admin page "Reloading (Global)" and the server is offline
* Fix issue where server uptime can sometimes display a bunch of extra data

## [3.1.5] 7/2/2023

* Add UI for managing static zone launcher `minZoneProcesses`
* When `minZoneProcesses` is set to 0, the Spire will set it to 10 by default. This is to prevent the server from booting with no zone processes

![image](https://github.com/EQEmu/Server/assets/3319450/63f6395b-b4c4-46eb-8fbe-aa6ea4ca6893)

## [3.1.4] 6/10/2023

Add command `spire:launcher [start|stop|restart]` to start, stop, or restart the launcher. This will be used in the future in tandem with installer v2 and utility scripts to manage the launcher

## [3.1.3] 4/30/2023

More improvements to pre-flight checks, they should fail pass-through now and not block the server from booting

## [3.1.1] 4/29/2023

* Add command `spire:occulus-update` to update Occulus to the latest version or install it if it doesn't exist [#116](https://github.com/Akkadius/spire/pull/116)

## [3.1.0] 4/29/2023

* Improvements to how Spire acquires static assets

## [3.0.8] 4/28/2023

* Spire will now exit immediately after updating if no terminal is attached [#114](https://github.com/Akkadius/spire/pull/114)
* Spire web interface port can now be configured using **eqemu_config.json** @ **spire.http_port** if it exists or environment variable `SPIRE_HTTP_PORT` if it exists [#112](https://github.com/Akkadius/spire/pull/114)
* Add command `spire:init [username] [password] --auth-enabled=true` to initialize a new Spire install with a default admin [#113](https://github.com/Akkadius/spire/pull/114)
* Add command `user:change-password [username] [password]` to change a user's password [#113](https://github.com/Akkadius/spire/pull/114)

## [3.0.7] 4/15/2023 - Spire Admin

* Added an asterisk to optional services in the pre-flight checks to make it more clear that they are optional
* Fixed an issue with pre-flight checks where an optional service would keep the operator from booting the server
* Fixed an issue where pre-flight checks would not clear between runs giving false reporting on checks

![image](https://user-images.githubusercontent.com/3319450/232274332-a969242a-45ac-4da1-a284-b8f0854a742e.png)

## [3.0.6] 4/12/2023

Fix issue with NPC/Bots editors that were incorrectly sending types on save in the edit form 

## [3.0.5] 4/4/2023 - Spire Admin

Fix issue with chat message returns bleeding into other telnet API command output

## [3.0.4] 4/3/2023 - Spire Admin

Fix a regression with Spells and Items editor where changing tabs would wipe out edits made in other tabs

## [3.0.3] 4/3/2023 - Spire Admin

Fix an issue during Spire initialization with older MySQL installs and a collation, index size issue

## [3.0.2] 4/3/2023 - Spire Admin

* Fixed an issue where when the control key is pressed while in an input field that it would bring up a modal help window. This is annoying in circumstances where you use "Ctrl+A" to select all in a text field to delete and the modal pops up. The modal is now cancelled under this scenario
* Fixed an issue where HTTP streams are incorrectly sending `Bearer` headers when a token is not set
* JWT signing key will now use the encryption key
* When local installs have authentication disabled, local JWT tokens get purged so API calls aren't made and authentication is rejected

## [3.0.1] 4/3/2023 - Spire Admin

Fix issue with Windows resolving `localhost` to ipv6 for server API connections

## [3.0.0] 4/2/2023 - Spire Admin

This release is nothing short of massive. Buckle up! The changelog notes are hefty!

### [Spire Admin] Dashboard

Similar to the dashboard in [Occulus](https://github.com/Akkadius/Occulus), but ported for modernization, quality of life changes, and the high-level server information is always available at the top of every admin page.

![image](https://user-images.githubusercontent.com/3319450/229333545-e8c6132c-68f7-4971-8d7c-0e00eda67eb6.png)

### [Spire Admin] Static Zone Launcher

You can now manage launch static zones via the Occulus launcher directly in the server start UI.

![image](https://user-images.githubusercontent.com/3319450/229336435-f2f1a173-b1f9-4c6e-b53b-1b79f19af9bc.png)

### [Spire Admin] Server Pre-Flight Checks

Before you launch your server, you are required to run basic pre-flight checks, this will ensure that your processes don't have any issues and you are made aware of them before all of your processes are launched in the background

![image](https://user-images.githubusercontent.com/3319450/229333233-b0fc22b1-0df0-467b-8b13-53621c382306.png)

### [Spire Admin] Players Online

![image](https://user-images.githubusercontent.com/3319450/229333775-104344eb-e008-48a6-9acd-13a340401bdc.png)

### [Spire Admin] Zone Servers

List and search booted zone processes and resources, kill zone processes

![image](https://user-images.githubusercontent.com/3319450/229333800-b39a8443-71fb-49c6-8af8-54d9ba0391dd.png)

### [Spire Admin] Manual Backups

Manually backup server assets

![image](https://user-images.githubusercontent.com/3319450/229333864-de528f46-c81a-4404-9327-cb06dd2bdddc.png)

### [Spire Admin] EverQuest Client Asset Exports (Publicly Downloadable Links)

![image](https://user-images.githubusercontent.com/3319450/229333921-73b79f4c-d1d6-44f8-bdee-bead89b847b5.png)

### [Spire Admin] Server Configuration Management

Tired of editing that pesky `eqemu_config` JSON file manually? Spire admin breaks out configuration editing into easy to use deep-link able sections

![image](https://user-images.githubusercontent.com/3319450/229333958-4eb2b9de-38d0-43d0-8dd9-67bd5c666a19.png)

### [Spire Admin] Crash Log Webhooks (Occulus)

Crash log webhooks that you used in Occulus? Spire integrates with Occulus to continue providing those in Spire. You can edit the Discord webhook right in the UI

![image](https://user-images.githubusercontent.com/3319450/229334006-10d5e97f-09c9-4f77-8bfd-98552ca90d60.png)

### [Spire Admin] Message of the Day

You know that message players read when they log in? You edit that here

![image](https://user-images.githubusercontent.com/3319450/229334040-c8a42200-53ec-4dc1-83e9-91a5a80d08cd.png)

### [Spire Admin] Quest Hot Reload (Occulus)

The hot reload functionality that was brought to you in Occulus is now configurable directly from Spire Admin. Anytime you make changes here it will immediately take affect.

Critical bug fixes have also been made to the HRM

![image](https://user-images.githubusercontent.com/3319450/229334070-96250391-8cc4-4c9f-b00c-e267d0f96969.png)

### [Spire Admin] Server Reloading

You can trigger server reloads directly through Spire. Reloads are triggered globally and reload data from the database into the server memory state.

Many tools already utilize this when editing related data.

![image](https://user-images.githubusercontent.com/3319450/229336522-158c268e-48bb-4180-8ca3-e53a9bc4e446.png)

### [Spire Admin] Server Rules

List, search, edit your server rules in real time. Spire will reload rules directly on your server immediately so you do not need to manually reload them in game. Super handy!

![image](https://user-images.githubusercontent.com/3319450/229334145-4536565e-50b2-4b17-bc7a-7f0bf485bd92.png)

### [Spire Admin] Server Database Backup

Ever just want to take a quick manual database backup? Spire admin could not make this any easier by tapping into the native eqemu server utility that wraps around `mysqldump`. You can select specific types of tables to dump, dump everything and compress it with an immediate download

![image](https://user-images.githubusercontent.com/3319450/229334188-57e515d1-efee-4572-9b6d-620b4d1cbb87.png)


### [Spire Admin] Server File Log Viewer

Completely overhauled and ported from Occulus. Spire file log viewer is highly performant, streams realtime and only streams changes since the last file write.

Emulates console coloring, filters by log type, search filters, file deletion, bulk file deletion and more!

<video src="https://user-images.githubusercontent.com/3319450/229334302-064df32f-9f27-405a-a6f0-c5623ab35400.mp4" data-canonical-src="https://user-images.githubusercontent.com/3319450/229334302-064df32f-9f27-405a-a6f0-c5623ab35400.mp4" controls="controls" muted="muted" class="d-block rounded-bottom-2 border-top width-fit" style="max-height:640px; min-height: 200px"></video>

### [Spire Admin] Discord Logging Webhooks

Manage your Discord logging webhooks directly through Spire.

![image](https://user-images.githubusercontent.com/3319450/229334410-111c2cc4-40d6-4644-8426-55f7e70ef64d.png)

### [Spire Admin] Manage Server Logging Settings

Manage your logging settings directly in Spire admin. Anytime changes are made your logging settings are reloaded real-time in the server without having to reboot or type commands in game!

![image](https://user-images.githubusercontent.com/3319450/229334469-85521c13-5435-4419-b9c2-2102949d7043.png)

### [Spire Admin] Player Event Log Explorer

Accompanied by a massive new feature in the EverQuest Emulator server ([see Player Event Logging](https://github.com/EQEmu/Server/pull/2833)) we are brought a powerful new explorer for rich auditing tools.

#### Rich Data Viewing

Spire translates event data into viewable NPC, Item, Spell etc. cards to have contextual insight into the events and what they contain.

![image](https://user-images.githubusercontent.com/3319450/216812922-547efee0-4a84-4b1d-8637-1a2d13283cba.png)

![image](https://user-images.githubusercontent.com/3319450/216810318-32e0e8f3-1392-497e-a14b-55f6585e7edd.png)

#### Event Raw

Any event you can see the richly formatted event or you can inspect the raw event data

![event-raw](https://user-images.githubusercontent.com/3319450/216810779-b14bbb99-7492-497e-afdf-407843ab1275.gif)

#### Flexible Filtering

You can filter by event type, zone_id, character_id which are top level filters always available and indexed regardless of event type

![event-filter](https://user-images.githubusercontent.com/3319450/216810442-62040911-f17c-4b41-b032-ec00a440c49d.gif)

#### Advanced Event Data Filtering

You can filter by JSON event data

![event-data-filter](https://user-images.githubusercontent.com/3319450/216810590-9802cfb5-a619-4e08-ab0e-342376e4044b.gif)

You can even filter by deeply nested data. For example if you're trying to search for a certain item that was traded by doing a nested wildcard search

![event-deep-filter](https://user-images.githubusercontent.com/3319450/216810644-dea5d9b6-3a8c-4936-afaf-efae03fb1fe5.gif)

![image](https://user-images.githubusercontent.com/3319450/229334536-f035ccd4-45ae-48e0-b5f7-e5b57a7b80e0.png)

### [Spire Admin] Player Event Log Settings

Once again you're able to edit your player event log settings within spire admin and changes are real-time reloaded in game without having to issue commands!

![image](https://user-images.githubusercontent.com/3319450/229334708-ab18f1d2-d9bb-49c8-8689-747bce0c2453.png)

### [Spire Admin] Server Update

Spire admin now has a rich interface where you can download new server releases, or build your own (Linux only)

![image](https://user-images.githubusercontent.com/3319450/229334953-b945783c-9e6a-4356-90d8-3a53d0dc7f21.png)

![image](https://user-images.githubusercontent.com/3319450/229334987-c7ac5374-a4ca-436e-b52f-9d8b7edf471f.png)

![image](https://user-images.githubusercontent.com/3319450/229334994-e5ca8f0f-6f65-442d-90f2-691438792c30.png)

![image](https://user-images.githubusercontent.com/3319450/229334970-cae8036a-5d15-44ce-87f5-67bcc9fddb09.png)

### [Spire Admin] Permissions

If you'd like to limit access of other users to your server and especially sensitive areas. There are permissions that cover access to various resources within the admin section of Spire in your local installation

![image](https://user-images.githubusercontent.com/3319450/229336838-8e7a73d7-8e05-4737-a89f-1a61e0cae0e1.png)

![image](https://user-images.githubusercontent.com/3319450/229336854-19a2b2d8-8937-497f-ae47-57ba21d27e22.png)

![image](https://user-images.githubusercontent.com/3319450/229336858-d0becf18-f67f-45aa-b543-1db54685e774.png)

![image](https://user-images.githubusercontent.com/3319450/229336867-70706d23-d286-4ada-a04a-039f0a6903fe.png)

### Spire Setup

Spire when ran locally will now redirect users to a one-time initial setup where users can choose whether they want their install to be guarded by authentication or not.

This bootstraps Spire for local installations.

![image](https://user-images.githubusercontent.com/3319450/229335122-9ee573db-93e4-48b5-8c8a-b836c0bd03cf.png)

### Local Users

Spire now has the ability during initial install to be configured with local users and authentication!

![image](https://user-images.githubusercontent.com/3319450/229334852-bc03242f-fe94-41f2-9a97-3ab866cdcf90.png)

![image](https://user-images.githubusercontent.com/3319450/229334891-1fba993e-d298-4fee-9234-479c14db4e8c.png)


### Navigation Search

Want to quickly and fluidly navigate your way around Spire? There's now a kbar in Spire similar to Discord, Slack, Github etc.

Navigation search can be activated by using Ctrl + K, Ctrl + / or clicking "Search" in the navigation menu

![image](https://user-images.githubusercontent.com/3319450/229332345-76c89678-ccc5-4850-b223-39dd0af5f712.png)

### Keyboard Shortcuts

My pressing the **Control** key, you can now see a modal that will show you Spire keyboard shortcuts

![image](https://user-images.githubusercontent.com/3319450/229333312-8fd5fd05-6fa2-42bc-8fc5-27aef60b1f6f.png)

### Item Editor

* Item editor edit button now is a hyperlink that can be opened in a new tab https://github.com/Akkadius/spire/issues/106
* Item editor clone button now is a hyperlink that can be opened in a new tab https://github.com/Akkadius/spire/issues/106
* Item editor now can deeply link to tabs
* Added "Food" tab with proper item display values in item card preview https://github.com/Akkadius/spire/issues/100
* Added proper "zero value" for skillmodtype and extradmgskill (Thanks Kinglykrab)

### Spell Editor

* Add "Corruption" to resist types https://github.com/Akkadius/spire/issues/92
* Spell editor now can deeply link to tabs
* Have DB String buttons open a new tab to prevent loss of progress in editor https://github.com/Akkadius/spire/issues/95

### Merchant Editor

* Fix an issue with new merchantlist entries creating a max_status of 0 preventing new merchant items from being displayed https://github.com/Akkadius/spire/issues/107

### NPC / Bot Spells Editor

* Fix an issue where editing a spell results in a number type error https://github.com/Akkadius/spire/issues/83

### Quest API Explorer

* Fix an issue with copy to clipboard for multiple events being displayed in search https://github.com/Akkadius/spire/issues/104
* Added user preference default language https://github.com/Akkadius/spire/issues/98

### Fixes

* Fixed [Bug] Merchant Editor creates new merchant list entries with max status of 0 https://github.com/Akkadius/spire/issues/107

## [2.6.1]

### Item Editor

- Fix an issue where charm file ID would not save properly when numbers were used in the field

## [2.6.0]

### Auto Update Tweaks

- Added a request timeout to the automatic updater routine

## [2.5.9]

### Windows Manifest Fix

- Fix an issue where the executable was getting flagged as suspicious

## [2.5.6]

### Client Files Export Fix

- Fix an issue where when two or more databases are present client files export will export multiples of the same column

## [2.5.5]

### Item Card Display

- Fix an issue where mods similar to "combat effects" were rendering on the preview card even if the value was 0

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
