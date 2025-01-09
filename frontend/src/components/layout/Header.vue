<template>
  <div class="header py-4">
    <div class="container">
      <div class="d-flex">
        <a class="header-brand" href="/">
          <img src="~@/assets/img/eqemu-logo.png" class="header-brand-img">
          Server Admin
          <small class="server-longname mobile-hide" style="margin-left:10px; font-size: 16px;"></small>
        </a>
        <div class="d-flex order-lg-2 ml-auto">

          <!-- Night mode toggle -->
          <div class="nav-item d-none d-md-flex">
            <label class="custom-switch" style="display:none">
              <i class="fe fe-sun" style="margin-right:10px"></i>
              <input type="checkbox" name="custom-switch-checkbox" class="custom-switch-input night-mode-toggle"
                     onclick="toggleDarkMode()">
              <span class="custom-switch-indicator"></span>
            </label>
          </div>

          <!-- EQEmu Github -->
          <div class="nav-item d-none d-md-flex">
            <a href="https://github.com/EQEmu/Server"
               class="btn btn-outline-primary btn-sm"
               data-original-title="Github"
               data-toggle="tooltip"
               target="eqemu-github">
              <i class="fe fe-github" style="font-size:12px"></i>
              EQEmu/Server
            </a>
          </div>

          <div class="dropdown">
            <a href="#" class="nav-link pr-0 leading-none" data-toggle="dropdown">
              <span class="avatar avatar-image"></span>
              <span class="ml-2 d-none d-lg-block">
                                  <span class="text-default"> {{ username }}</span>
                <!-- <small class="text-muted d-block mt-1"> {{ username }}</small> -->
                                </span>
            </a>
            <div class="dropdown-menu dropdown-menu-right dropdown-menu-arrow">

              <a class="dropdown-item" href="/logout">
                <i class="dropdown-icon fe fe-log-out"></i> Sign out
              </a>
            </div>
          </div>
        </div>
        <a href="#" class="header-toggler d-lg-none ml-3 ml-lg-0" data-toggle="collapse"
           data-target="#headerMenuCollapse">
          <span class="header-toggler-icon"></span>
        </a>
      </div>
    </div>
  </div>
</template>

<script>

  export default {
    data() {
      return {
        username: ''
      }
    },
    created() {
      // const accessTokenUserInfo = ApiClient.getAccessTokenUserInfo()
      // this.username             = accessTokenUserInfo.user
    }
  }
</script>
