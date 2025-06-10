<template>
  <content-area class="text-center login-center-container fade-in">

    <div class="row justify-content-center">
      <div class="col-lg-3 col-sm-12 col-md-12">

        <router-link class="ml-3 mt-3 mb-3" :to="ROUTE.HOME">
          <h1
            style="font-size: 100px"
            class="text-center eq-header mb-0"
          >
            Spire
          </h1>

        </router-link>

        <eq-window class="fade-in login-box" v-if="hasAuthOptions()">
            <form
              class="m-0 p-0"
              style="top:50%;"
              v-on:submit.prevent="loginSpire"
            >

              <!-- Local Login -->
              <div v-if="localAuthEnabled" class="fade-in">
                <div class="form-group">
                  <label for="formGroupExampleInput">Username</label>
                  <input
                    type="text"
                    class="form-control"
                    autocomplete="off"
                    id="username"
                    name="username"
                    v-model="username"
                    placeholder="Your username..."
                  >
                </div>
                <div class="form-group">
                  <label for="formGroupExampleInput2">Password</label>
                  <div class="input-group">
                    <input
                      :type="(showPassword ? 'text' : 'password')"
                      class="form-control"
                      autocomplete="off"
                      id="password"
                      name="password"
                      v-model="password"
                      v-on:keyup.enter="loginSpire"
                      placeholder="Your password..."
                    >

                    <div class="input-group-append">
                      <button class="btn btn-dark" type="button" @click="showPassword = !showPassword">
                        <i class="fe fe-lock"></i>
                        Show Password
                      </button>
                    </div>
                  </div>

                </div>

                <button
                  type="submit"
                  class="btn btn-lg btn-dark btn-block"
                  @click="loginSpire()"
                  style="color:white"
                >
                  <i class="fe fe-lock"></i>
                  Spire Login
                </button>
              </div>

              <a
                class="btn btn-lg btn-dark btn-block"
                @click="loginGithub()"
                style="color:white"
                v-if="githubAuthEnabled"
              >
                <i class="fe fe-github"></i>
                Sign in with Github
              </a>


            </form>

            <info-error-banner
              :slim="true"
              :notification="notification"
              :error="error"
              @dismiss-error="error = ''"
              @dismiss-notification="notification = ''"
              class="mt-3"
            />
        </eq-window>


      </div>
    </div>
  </content-area>
</template>

<script>
import ContentArea     from "../components/layout/ContentArea";
import {AppEnv}        from "../app/env/app-env";
import {ROUTE}         from "../routes";
import {SpireApi}      from "../app/api/spire-api";
import InfoErrorBanner from "@/components/InfoErrorBanner.vue";
import UserContext     from "@/app/user/UserContext";
import axios           from "axios";
import EqWindow from "@/components/eq-ui/EQWindow.vue";

export default {
  name: 'Login.vue',
  components: {EqWindow, InfoErrorBanner, ContentArea },
  data() {
    return {

      // login
      username: "",
      password: "",

      showPassword: false,

      // api responses
      error: "",
      notification: "",

      // auth mechanisms
      localAuthEnabled: false,
      githubAuthEnabled: false,

      // route constants
      ROUTE: ROUTE,
    }
  },
  async mounted() {
    await AppEnv.init()

    // auth settings
    this.localAuthEnabled  = AppEnv.isLocalAuthEnabled()
    this.githubAuthEnabled = AppEnv.isGithubAuthEnabled()

    // check query params user,password,redirect
    this.$router.onReady(async () => {
      const query = this.$route.query
      if (query && query.user && query.password) {
        console.log("login with query params")
        this.username = query.user
        this.password = query.password
        await this.loginSpire()
      }
    })
  },
  methods: {
    async loginSpire() {
      try {
        const r = await axios.post(SpireApi.getBasePath() + "/auth/login", {
          username: this.username,
          password: this.password,
        })
        if (r.data && r.status === 200) {
          this.notification = "Login succeeded!"

          if (r.data && r.data.data && r.data.data.token) {
            UserContext.storeAccessToken(r.data.data.token)
            SpireApi.reloadAxios()

            setTimeout(() => {
              const query = this.$route.query
              if (query.redirect) {
                this.$router.push(query.redirect).catch((e) => {
                })
              } else {
                this.$router.push(ROUTE.HOME).catch((e) => {
                })
              }

            }, 100)
          }
        }
      } catch (err) {
        // error notify
        if (err.response && err.response.data && err.response.data.error) {
          this.error = err.response.data.error
        }
      }
    },

    hasAuthOptions() {
      return this.githubAuthEnabled || this.localAuthEnabled
    },
    loginGithub: function () {
      const width  = 800
      const height = 800
      const left   = (screen.width / 2) - (width / 2);
      const top    = (screen.height / 2) - (height / 2);
      const url    = SpireApi.getBasePath() + '/auth/github';
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
  padding-top: 40px;
  padding-bottom: 40px;
  background-color: #f5f5f5;
}

.form-signin {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: auto;
}

.form-signin .checkbox {
  font-weight: 400;
}

.form-signin .form-control {
  position: relative;
  box-sizing: border-box;
  height: auto;
  padding: 10px;
  font-size: 16px;
}

.form-signin .form-control:focus {
  z-index: 2;
}

.form-signin input[type="email"] {
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}

.form-signin input[type="password"] {
  margin-bottom: 10px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}

.bd-placeholder-img {
  font-size: 1.125rem;
  text-anchor: middle;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

@media (min-width: 768px) {
  .bd-placeholder-img-lg {
    font-size: 3.5rem;
  }
}


.login-center-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  min-height: 100vh;
}

</style>
