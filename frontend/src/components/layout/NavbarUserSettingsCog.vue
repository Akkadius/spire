<template>
  <div>
    <a href="#sidebarModalActivity" class="navbar-user-link" data-toggle="modal" v-b-modal.user-settings-modal>
      <span class="icon">
        <i class="fe fe-settings"></i>
      </span>
    </a>

    <b-modal id="user-settings-modal" centered title="Settings" size="lg">

      <!-- Debug Mode -->
      <div class="row mb-4">
        <div class="col-4 text-right">
          Debug Mode
          <b-form-checkbox
            v-model="debugEnabled"
            name="check-button"
            @change="debugUpdate"
            switch
            class="d-inline-block ml-3"
          />
        </div>
        <div class="col-8">
          <small class="text-muted">
            Some debugging features may require browser reload to take affect
          </small>
        </div>
      </div>

      <!-- Tab hovering -->
      <div class="row mb-4">
        <div class="col-4 text-right">
          Enable Tab Hover
          <b-form-checkbox
            v-model="tabHoverModeEnabled"
            name="check-button"
            @change="updateSetting('tab-hover', tabHoverModeEnabled)"
            switch
            class="d-inline-block ml-3"
          />
        </div>
        <div class="col-8">
          <small class="text-muted">
            Tabs are activated via mouse hover versus click
          </small>
        </div>
      </div>

      <!-- Alpha -->
      <div class="row mb-4">
        <div class="col-4 text-right">
          Enable Alpha Tools
          <b-form-checkbox
            v-model="alphaToolsEnabled"
            name="check-button"
            @change="alphaUpdate"
            switch
            class="d-inline-block ml-3"
          />
        </div>
        <div class="col-8">
          <small class="text-muted">
            Tools in alpha are enabled
          </small>
          <span class="badge badge-primary">ALPHA</span>

        </div>
      </div>

      <!-- Spell Legacy Icons -->
      <div class="row mb-4">
        <div class="col-4 text-right">
          Spell Legacy Icons
          <b-form-checkbox
            v-model="spellLegacyIcons"
            name="check-button"
            @change="spellLegacyIconsUpdate"
            switch
            class="d-inline-block ml-3"
          />
        </div>
        <div class="col-8">
          <small class="text-muted">
            Enables legacy spell icons
          </small>
        </div>
      </div>

      <template #modal-footer>
        <div class="">

        </div>
      </template>
    </b-modal>
  </div>
</template>

<script>

import {LocalSettings, Setting} from "@/app/local-settings/localsettings";
import {App}                    from "@/constants/app";
import {EventBus}               from "@/app/event-bus/event-bus";

export default {
  name: "NavbarUserSettingsCog",
  data() {
    return {
      debugEnabled: LocalSettings.isDebugEnabled(),
      tabHoverModeEnabled: LocalSettings.isTabHoverEnabled(),
      alphaToolsEnabled: LocalSettings.isAlphaEnabled(),
      spellLegacyIcons: LocalSettings.isSpellLegacyIconsEnabled(),
    }
  },
  methods: {
    debugUpdate() {
      // checkbox apparently hasn't had enough time to update reactively... queue it
      setTimeout(() => {
        LocalSettings.set(Setting.DEBUG_MODE, this.debugEnabled)
        App.DEBUG = this.debugEnabled
        EventBus.$emit('DEBUG_UPDATED', true);
      }, 100)
    },
    alphaUpdate() {
      // checkbox apparently hasn't had enough time to update reactively... queue it
      setTimeout(() => {
        LocalSettings.set(Setting.ALPHA_ENABLED, this.alphaToolsEnabled)
        App.ALPHA_TOOLS_ENABLED = this.alphaToolsEnabled
        EventBus.$emit('ALPHA_ENABLED', true);
      }, 10)
    },
    spellLegacyIconsUpdate() {
      // checkbox apparently hasn't had enough time to update reactively... queue it
      setTimeout(() => {
        LocalSettings.set(Setting.SPELL_LEGACY_ICONS, this.spellLegacyIcons)
        App.SPELL_LEGACY_ICONS_ENABLED = this.spellLegacyIcons
        EventBus.$emit('SPELL_LEGACY_ICONS_ENABLED', true);
      }, 10)
    },
    updateSetting(name, value) {
      LocalSettings.set(name, value)
    }
  }


}
</script>
