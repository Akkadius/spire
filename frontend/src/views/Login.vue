<template>
  <content-area class="text-center centered">
    <div class="row justify-content-center">
      <div class="col-3">

        <router-link class="ml-3 mt-3 mb-3" to="/">
          <h1 class="text-center eq-header small-mobile">
            Spire
          </h1>
        </router-link>

        <div class="card">
          <div class="card-body">
            <form class="form-signin" style="top:50%;">

              <a class="btn btn-lg btn-dark btn-block" @click="loginGithub()" style="color:white">
                <i class="fe fe-github"></i>
                Sign in with Github
              </a>

            </form>
          </div>
        </div>

      </div>
    </div>
  </content-area>
</template>

<script>
import ContentArea from "../components/layout/ContentArea";
import {AppEnv}    from "../app/env/app-env";
export default {
  name: 'Login.vue',
  components: { ContentArea },
  data() {
    return {
      githubAuthEnabled: AppEnv.isGithubAuthEnabled(),
    }
  },
  mounted() {
    setTimeout(() => {
      this.githubAuthEnabled = AppEnv.isGithubAuthEnabled()
    }, 1000)
  },
  methods: {
    loginGithub: function () {
      const width  = 800
      const height = 800
      const left   = (screen.width / 2) - (width / 2);
      const top    = (screen.height / 2) - (height / 2);
      const url    = process.env.VUE_APP_BACKEND_BASE_URL + '/auth/github';
      const title  = 'Github'
      const win    = window.open(url, title, 'toolbar=no, location=no, directories=no, status=no, menubar=no, scrollbars=no, resizable=no, copyhistory=no, width=' + width + ', height=' + height + ', top=' + top + ', left=' + left);
      var timer    = setInterval(() => {
        if (win.closed) {
          clearInterval(timer);
          this.$router.push({ path: '/' })
        }
      }, 100);
    }
  }
}
</script>

<style scoped>
html,
body {
  height: 100%;
}
body {
  padding-top:      40px;
  padding-bottom:   40px;
  background-color: #f5f5f5;
}
.centered {
  position:          fixed;
  top:               50%;
  left:              50%;
  transform:         translate(-50%, -50%);
  -webkit-transform: translate(-50%, -50%);
  -moz-transform:    translate(-50%, -50%);
  -o-transform:      translate(-50%, -50%);
  -ms-transform:     translate(-50%, -50%);
  font-size:         20px;
  padding:           5px;
  z-index:           100;
}
.form-signin {
  width:     100%;
  max-width: 330px;
  padding:   15px;
  margin:    auto;
}
.form-signin .checkbox {
  font-weight: 400;
}
.form-signin .form-control {
  position:   relative;
  box-sizing: border-box;
  height:     auto;
  padding:    10px;
  font-size:  16px;
}
.form-signin .form-control:focus {
  z-index: 2;
}
.form-signin input[type="email"] {
  margin-bottom:              -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius:  0;
}
.form-signin input[type="password"] {
  margin-bottom:           10px;
  border-top-left-radius:  0;
  border-top-right-radius: 0;
}
.bd-placeholder-img {
  font-size:           1.125rem;
  text-anchor:         middle;
  -webkit-user-select: none;
  -moz-user-select:    none;
  -ms-user-select:     none;
  user-select:         none;
}
@media (min-width: 768px) {
  .bd-placeholder-img-lg {
    font-size: 3.5rem;
  }
}
</style>
