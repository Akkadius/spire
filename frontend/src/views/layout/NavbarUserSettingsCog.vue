<template>
  <div>
    <a href="#sidebarModalActivity" class="navbar-user-link" data-toggle="modal" v-b-modal.user-settings-modal>
              <span class="icon">
                <i class="fe fe-settings"></i>
              </span>
    </a>

    <b-modal id="user-settings-modal" centered title="Settings" size="lg">

      <b-form-checkbox v-model="debugEnabled" name="check-button" @change="debugUpdate" switch>
        Debug Mode <small class="text-muted pl-5">Some debugging features may require browser reload to take
        affect</small>
      </b-form-checkbox>

      <template #modal-footer>
        <div class="">

        </div>
      </template>
    </b-modal>
  </div>
</template>

<script>

import LocalSettings from "@/app/local-settings/localsettings";
import {App} from "@/constants/app";
import {EventBus} from "@/app/event-bus/event-bus";

export default {
  name: "NavbarUserSettingsCog",
  data() {
    return {
      debugEnabled: LocalSettings.get("debug-mode") ? LocalSettings.get("debug-mode") : false
    }
  },
  methods: {
    debugUpdate() {
      // checkbox apparently hasn't had enough time to update reactively... queue it
      setTimeout(() => {
        LocalSettings.set("debug-mode", this.debugEnabled)
        App.DEBUG = this.debugEnabled
        EventBus.$emit('DEBUG_UPDATED', true);
      }, 100)

    }
  }


}
</script>
