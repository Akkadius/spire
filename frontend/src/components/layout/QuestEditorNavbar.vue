<template>
  <nav
    class="navbar navbar-vertical fixed-left navbar-expand-md navbar-dark navbar-vibrant pb-0"
    id="sidebar">
    <div class="container-fluid">

      <!-- Toggler -->
      <button
        class="navbar-toggler" type="button" data-toggle="collapse" data-target="#sidebarCollapse"
        aria-controls="sidebarCollapse" aria-expanded="false" aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>

      <router-link class="ml-3 mt-3" to="/">
        <h1 class="text-center eq-header small-mobile">
          Spire
        </h1>

      </router-link>


      <div class="collapse navbar-collapse p-0" id="sidebarCollapse">
        <sl-vue-tree
          v-model="nodes"
          ref="slVueTree"
          class="mt-3"
          style="width: 225px; height: 100%; margin: 0;"
        >
          <template slot="title" slot-scope="{ node }">

          <span class="item-icon">
            <i class="fa fa-file" v-if="node.isLeaf"></i>
            <i class="fa fa-folder" v-if="!node.isLeaf"></i>
          </span>

          {{ node.title }}

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
import SlVueTree             from "sl-vue-tree"
// import "node_modules/sl-vue-tree/sl-vue-tree-dark.css"

export default {
  name: "QuestEditorNavbar",
  components: { NavSectionComponent, NavbarDropdownMenu, NavbarUserSettingsCog, SlVueTree },
  data() {
    return {
      nodes: [
        { title: 'Item1', isLeaf: true },
        { title: 'Item2', isLeaf: true, data: { visible: false } },
        { title: 'Folder1' },
        {
          title: 'Folder2', isExpanded: true, children: [
            { title: 'Item3', isLeaf: true },
            { title: 'Item4', isLeaf: true }
          ]
        },
        { title: 'Item1', isLeaf: true },
        { title: 'Item2', isLeaf: true, data: { visible: false } },
        { title: 'Folder1' },
        {
          title: 'Folder2', isExpanded: true, children: [
            { title: 'Item3', isLeaf: true },
            { title: 'Item4', isLeaf: true }
          ]
        },
        { title: 'Item1', isLeaf: true },
        { title: 'Item2', isLeaf: true, data: { visible: false } },
        { title: 'Folder1' },
        {
          title: 'Folder2', isExpanded: true, children: [
            { title: 'Item3', isLeaf: true },
            { title: 'Item4', isLeaf: true }
          ]
        },
        { title: 'Item1', isLeaf: true },
        { title: 'Item2', isLeaf: true, data: { visible: false } },
        { title: 'Folder1' },
        {
          title: 'Folder2', isExpanded: true, children: [
            { title: 'Item3', isLeaf: true },
            { title: 'Item4', isLeaf: true }
          ]
        },
        { title: 'Item1', isLeaf: true },
        { title: 'Item2', isLeaf: true, data: { visible: false } },
        { title: 'Folder1' },
        {
          title: 'Folder2', isExpanded: true, children: [
            { title: 'Item3', isLeaf: true },
            { title: 'Item4', isLeaf: true }
          ]
        },
        { title: 'Item1', isLeaf: true },
        { title: 'Item2', isLeaf: true, data: { visible: false } },
        { title: 'Folder1' },
        {
          title: 'Folder2', isExpanded: true, children: [
            { title: 'Item3', isLeaf: true },
            { title: 'Item4', isLeaf: true }
          ]
        },
        { title: 'Item1', isLeaf: true },
        { title: 'Item2', isLeaf: true, data: { visible: false } },
        { title: 'Folder1' },
        {
          title: 'Folder2', isExpanded: true, children: [
            { title: 'Item3', isLeaf: true },
            { title: 'Item4', isLeaf: true }
          ]
        },
        { title: 'Item1', isLeaf: true },
        { title: 'Item2', isLeaf: true, data: { visible: false } },
        { title: 'Folder1' },
        {
          title: 'Folder2', isExpanded: true, children: [
            { title: 'Item3', isLeaf: true },
            { title: 'Item4', isLeaf: true }
          ]
        },
      ]
    }
  },

  async mounted() {
    // expose instance to the global namespace for better debugging
    window.slVueTree = this.$refs.slVueTree;

    // window.addEventListener('keydown', (event) => this.onArrowDownHandler(event));
  },

  methods: {
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
        }
        else if (keyCode === 'ArrowLeft') {
          if (selectedNode.isLeaf) return;
          slVueTree.updateNode(selectedNode.path, { isExpanded: false });
        }
        else if (keyCode === 'ArrowRight') {
          if (selectedNode.isLeaf) return;
          slVueTree.updateNode(selectedNode.path, { isExpanded: true });
        }
        else if (keyCode === 'Enter' || keyCode === 'Space') {
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
