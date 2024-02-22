<template>
  <div class="sage-container" ref="sage-root">
    <iframe ref="sage-iframe" class="sage-iframe" src="/eqsage" title="EQSage"></iframe>
  </div>
</template>

<script type="ts">
import * as SpireApiTypes from "@/app/api";
import { SpireApi } from "../../app/api/spire-api";
import { SpireQueryBuilder } from "@/app/api/spire-query-builder";
import {Navbar}          from "../../app/navbar";
import { Zones } from "../../app/zones";
import { Spawn } from "../../app/spawn";
import { Npcs } from "../../app/npcs";
import { Grid } from "../../app/grid";
export default {
  components: {

  },
  data() {

    return {}
  },
  watch: {

  },

  destroyed() {
    Navbar.expand();
  },
  async mounted() {
    Navbar.collapse()
    this.$refs['sage-iframe'].addEventListener('load', e => {
      e.target.contentWindow.Spire = {
        SpireApi,
        SpireApiTypes,
        SpireQueryBuilder,
        Grid,
        Zones,
        Spawn,
        Npcs
      }
    })
  },
}
</script>

<style>
.sage-iframe {
  height: calc(100vh - 10px);
  border: none;
  margin-left: -35px;
  margin-top: -10px;
  width: calc(100% + 70px);
}
.code-display {
  font-size: 14px !important;
  max-width: 100% !important;
  width: 100%;
}
</style>
