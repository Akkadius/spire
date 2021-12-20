<template>
  <footer class="footer fixed-bottom" v-if="displayFooter">
    <div class="container">
      <div class="row align-items-center flex-row-reverse">
        <div class="col-auto ml-lg-auto">
          <div class="row align-items-center">
            <div class="col-auto">
              <ul class="list-inline list-inline-dots mb-0">
                <li class="list-inline-item">
                  <a href="https://github.com/EQEmu/Server/wiki" target="wiki">Documentation</a>
                </li>
              </ul>
            </div>
            <div class="col-auto">
              <a href="https://github.com/EQEmu/Server" class="btn btn-outline-primary btn-sm"
                 target="eqemu-github">
                <i class="fe fe-github" style="font-size:12px"></i>
                EQEmu/Server
              </a>
            </div>
          </div>
        </div>
        <div class="col-12 col-lg-6 mt-6 mt-lg-0 text-left">
          Everquest is a registered trademark of Daybreak Game Company LLC. <br>
          EQEmulator is not associated or affiliated in any way with Daybreak Game Company LLC.
        </div>
      </div>
    </div>
  </footer>

  <!-- Footer -->
  <!--
  <div class="footer">
      <div class="container">
          <div class="row">
              <div class="col-lg-8">
                  <div class="row">
                      <div class="col-6 col-md-3">
                          <ul class="list-unstyled mb-0">
                              <li><a href="#">First link</a></li>
                              <li><a href="#">Second link</a></li>
                          </ul>
                      </div>
                      <div class="col-6 col-md-3">
                          <ul class="list-unstyled mb-0">
                              <li><a href="#">Third link</a></li>
                              <li><a href="#">Fourth link</a></li>
                          </ul>
                      </div>
                  </div>
              </div>
              <div class="col-lg-4 mt-4 mt-lg-0">

              </div>
          </div>
      </div>
  </div>
  -->

  <!--
  <div class="card card-collapsed" style="position:fixed;top:100px;right:10px;width:400px">
      <div class="card-header">
          <small class="card-title"><i class="fe fe-power"></i> Server is rebooting in 10.4 minute(s)</small>
      </div>
  </div>
  -->
</template>

<script>
  export default {
    name: 'Footer.vue',
    data () {
      return {
        displayFooter: null
      }
    }
  }
</script>
