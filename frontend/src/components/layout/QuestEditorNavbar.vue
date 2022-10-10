<template>
  <nav
    class="navbar navbar-vertical fixed-left navbar-expand-md navbar-dark navbar-vibrant pb-0"
    id="sidebar"
  >
    <div class="container-fluid">

      <!-- Toggler -->
      <button
        class="navbar-toggler" type="button" data-toggle="collapse" data-target="#sidebarCollapse"
        aria-controls="sidebarCollapse" aria-expanded="false" aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>

      <router-link class="ml-3 mt-3" to="/">
        <h1 class="text-center eq-header small-mobile mb-0">
          Spire
        </h1>
        <span class="text-center eq-header" style="font-size: 25px;margin-left: 35px;">
          Quest Editor
        </span>
      </router-link>

      <div class="collapse navbar-collapse p-0" id="sidebarCollapse">
        <sl-vue-tree
          v-model="nodes"
          ref="slVueTree"
          :show-branches="false"
          class="mt-3"
          @select="nodeSelected"
          style="width: 405px; height: 88vh; margin: 0; background-color: rgba(20, 20, 20, 0.6); border: 1px solid rgb(30 30 30); overflow-x: scroll"
        >
          <template slot="title" slot-scope="{ node }">
          <span class="item-icon">
<!--            <i class="fa fa-file" v-if="node.isLeaf"></i>-->
            <i class="icon perl-icon medium-blue" v-if="node.isLeaf && node.title.includes('.pl')"></i>
            <i class="icon lua-icon medium-green" v-if="node.isLeaf && node.title.includes('.lua')"></i>
            <i class="fa fa-folder medium-yellow" v-if="!node.isLeaf"></i>
          </span>

            {{ node.title }}
          </template>

          <template slot="toggle" slot-scope="{ node }">
<!--          <span v-if="!node.isLeaf">-->
<!--            <i v-if="node.isExpanded" class="fa fa-chevron-down"></i>-->
<!--            <i v-if="!node.isExpanded" class="fa fa-chevron-right"></i>-->
<!--          </span>-->
          </template>
        </sl-vue-tree>
      </div>

    </div>
  </nav>
</template>

<script>

import NavbarDropdownMenu    from "@/components/layout/NavbarDropdownMenu";
import NavbarUserSettingsCog from "@/components/layout/NavbarUserSettingsCog";
import NavSectionComponent   from "@/components/layout/NavSectionComponent";
import SlVueTree  from "sl-vue-tree"
import {SpireApi} from "../../app/api/spire-api";
import path       from "path"
import * as util             from "util";
import {EventBus}            from "@/app/event-bus/event-bus";
// import "node_modules/sl-vue-tree/sl-vue-tree-dark.css"

export default {
  name: "QuestEditorNavbar",
  components: { NavSectionComponent, NavbarDropdownMenu, NavbarUserSettingsCog, SlVueTree },
  data() {
    return {
      files: null,
      loadedFolders: [""],
      nodes: [
        { title: 'Loading...', isLeaf: true },
      ]
    }
  },

  async mounted() {
    // expose instance to the global namespace for better debugging
    window.slVueTree = this.$refs.slVueTree;

    SpireApi.v1().get('/quest-file-api/list').then((response) => {
      if (response.data) {

        this.files = response.data.files

        this.loadTree();

        // console.log(questFiles)

      }
    });

    // window.addEventListener('keydown', (event) => this.onArrowDownHandler(event));
  },

  methods: {
    loadTree() {

      let questFiles = {}
      this.files.forEach((file) => {
        const filename = path.basename(file);
        const folder   = path.basename(path.dirname(file));

        if (folder === ".") {
          return false;
        }

        if (typeof questFiles[folder] === "undefined") {
          questFiles[folder] = [];
        }

        questFiles[folder].push(filename)
      })

      let nodes = []
      let folders = 0
      for (let folder in questFiles) {
        const files = questFiles[folder]

        // load children if the folder is loaded
        let children = []

        this.loadedFolders.forEach((f) => {
          if (f === folder) {
            files.forEach((file) => {
              let child      = {}
              child.title    = file
              child.isLeaf   = true
              child.fullPath = util.format("%s/%s", folder, file)
              children.push(child)
            })
          }
        })

        let node = {
          title: folder,
          isLeaf: false,
          isExpanded: false,
          children: children
        }

        nodes.push(node)
        folders++

      }

      this.nodes = nodes
    },

    nodeSelected(nodes, event) {

      // console.log(nodes[0])
      const slVueTree = this.$refs.slVueTree;

      let fullPath = []
      if (nodes[0]) {
        let fileName = nodes[0].title
        console.log(nodes[0])

        if (!nodes[0].isLeaf) {
          this.loadedFolders.indexOf(fileName) === -1 ? this.loadedFolders.push(fileName) : ""

          let questFiles = {}
          this.files.forEach((file) => {
            const filename = path.basename(file);
            const folder   = path.basename(path.dirname(file));

            if (folder === ".") {
              return false;
            }

            if (typeof questFiles[folder] === "undefined") {
              questFiles[folder] = [];
            }

            questFiles[folder].push(filename)
          })

          let children = []

          for (let folder in questFiles) {
            const files = questFiles[folder]
            if (folder === fileName) {
              files.forEach((file) => {
                let child      = {}
                child.title    = file
                child.isLeaf   = true
                child.fullPath = util.format("%s/%s", folder, file)
                children.push(child)
              })
            }
          }

          console.log(children)

          slVueTree.updateNode(nodes[0].path, { children: children });

          // this.loadedFolders.push(fileName)
          // this.loadTree()

          console.log("loaded folders %s", fileName)
        }

        if (nodes[0].path[0] >= 0) {
          let parentName = slVueTree.getNode([nodes[0].path[0]]).title

          console.log("p [%s] f [%s]", parentName, fileName)

          if (parentName !== fileName) {
            let fullFileName = util.format("%s/%s", parentName, fileName)

            // console.log(fullFileName)

            SpireApi.v1().get(util.format('/quest-file-api/%s', encodeURIComponent(fullFileName))).then((response) => {
              if (response.data) {
                // console.log(repsonse.data.contents)

                EventBus.$emit('EDITOR_OPEN_FILE', {
                  fileName: fileName,
                  fullFileName: fullFileName,
                  contents: response.data.content,
                });
              }
            });
          }
        }
      }
    },

    onArrowDownHandler(event) {
      event.preventDefault();
      const keyCode   = event.code;
      const slVueTree = this.$refs.slVueTree;

      if (slVueTree.selectionSize === 1) {
        const selectedNode = slVueTree.getSelected()[0];
        let nodeToSelect;

        if (keyCode === 'ArrowDown') {
          nodeToSelect = slVueTree.getNextNode(selectedNode.path, node => node.isVisible);
        } else if (keyCode === 'ArrowUp') {
          nodeToSelect = slVueTree.getPrevNode(selectedNode.path, node => node.isVisible);
        } else if (keyCode === 'ArrowLeft') {
          if (selectedNode.isLeaf) return;
          slVueTree.updateNode(selectedNode.path, { isExpanded: false });
        } else if (keyCode === 'ArrowRight') {
          if (selectedNode.isLeaf) return;
          slVueTree.updateNode(selectedNode.path, { isExpanded: false });
        } else if (keyCode === 'Enter' || keyCode === 'Space') {
          if (selectedNode.isLeaf) return;
          slVueTree.updateNode(selectedNode.path, { isExpanded: !selectedNode.isExpanded });
        }

        if (!nodeToSelect) return;

        slVueTree.select(nodeToSelect.path);

      } else if (keyCode === 'ArrowDown') {
        slVueTree.select(slVueTree.getFirstNode().path);
      } else if (keyCode === 'ArrowUp') {
        slVueTree.select(slVueTree.getLastNode().path);
      }
    }
  },
  watch: {
    $route(to, from) {
      this.hideNavbarAfterClick()
    }
  }
}
</script>

<style scoped>
.navbar-nav .nav-link > .fe {
  min-width: 1.25rem;
}

.navbar-dark.navbar-vibrant {
  background-image: linear-gradient(to bottom right, rgb(21 21 21 / 90%), rgb(191 126 70 / 90%)), url(~@/assets/img/eq-wallpaper-1.jpg);
}

@media only screen and (max-device-width: 640px) {
  .small-mobile {
    font-size: 40px !important;
  }
}

</style>
