<template>
  <content-area class="text-center fade-in">

    <div class="row justify-content-center mt-8">
      <div class="col-3">

        <router-link class="ml-3 mt-3 mb-3" :to="ROUTE.HOME">
          <h1
            style="font-size: 100px"
            class="text-center eq-header small-mobile">
            Spire
          </h1>
        </router-link>

        <div class="card" v-if="hasAuthOptions()">
          <div class="card-body">
            <form class="form-signin" style="top:50%;">

              <a class="btn btn-lg btn-dark btn-block" @click="loginGithub()" style="color:white" v-if="githubAuthEnabled">
                <i class="fe fe-github"></i>
                Sign in with Github
              </a>

            </form>
          </div>
        </div>

        <h2
          class="text-center eq-header small-mobile">
          Login
        </h2>

      </div>
    </div>
  </content-area>
</template>

<script>
import ContentArea      from "../components/layout/ContentArea";
import {AppEnv}         from "../app/env/app-env";
import {ROUTE}          from "../routes";
import {SpireApiClient} from "../app/api/spire-api-client";
export default {
  name: 'Login.vue',
  components: { ContentArea },
  data() {
    return {
      githubAuthEnabled: AppEnv.isGithubAuthEnabled(),
      ROUTE: ROUTE,
    }
  },
  mounted() {
    setTimeout(() => {
      this.githubAuthEnabled = AppEnv.isGithubAuthEnabled()
    }, 1000)
  },
  methods: {
    hasAuthOptions() {
      return this.githubAuthEnabled
    },
    loginGithub: function () {
      const width  = 800
      const height = 800
      const left   = (screen.width / 2) - (width / 2);
      const top    = (screen.height / 2) - (height / 2);
      const url    = SpireApiClient.getBasePath() + '/auth/github';
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
