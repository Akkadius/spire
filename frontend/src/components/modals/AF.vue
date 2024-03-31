<template>
  <b-modal
    id="af"
    centered
    no-fade
    no-close-on-backdrop
    no-close-on-esc
    title="Spire License Required"
    size="lg"
  >

    <template v-slot:modal-header-close>
      <div></div>
    </template>
    <div class="mt-3" v-if="!af">
      Dear Valued Spire User,<br><br>

      As we continue to strive for excellence in delivering unparalleled service to our community, we've encountered a
      rather unexpected challenge: our digital hamsters, the ones powering the very core of our software, have developed
      an insatiable appetite for premium, organic digital pellets. 🐹💻<br><br>

      <b>The Issue at Hand</b>: Our team, dedicated to sustainability and ethical digital hamster care, has seen
      operational costs skyrocket. The digital hamsters demand not just any pellets, but those harvested from the
      elusive digital fields of Silicon Valley, known for their high byte-nutrient content.<br><br>

      Additionally, the costs of maintaining our cloud infrastructure, now floating actual clouds for a more literal
      "cloud computing" experience, have simply become astronomical. 🌩️<br><br>

      <b>The Solution</b>: In light of these unmaintainable costs, we're excited to announce a shift to a premium
      subscription
      model, effective immediately! For the low, low price of just <b>$99.99 per month</b>, you'll gain exclusive access
      to:<br><br>

      <b>Unlimited Use</b>: Enjoy uninterrupted access to our software, now powered by digital hamsters running in
      diamond-encrusted wheels for that extra performance boost.<br><br>

      <b>Premium Support</b>: Direct access to our support team, who will now respond to inquiries via carrier pigeon
      for that personal touch. 🕊️<br><br>

      <b>Exclusive Features</b>: Get early access to features we're considering maybe possibly thinking about developing
      in the future.<br><br>

      <b>Special Offer</b>: Sign up in the next 24 hours to lock in your subscription at a 0.01% discount, and receive a
      complimentary digital hamster NFT (Non-Fungible Token), which is just as real as this entire announcement!<br><br>


      <b>Thank you for your continued support!</b><br><br>

      <b>Checkout with PayPal to continue using Spire.</b><br><br>
    </div>

    <div class="text-center" v-if="!af">
      <img
        src="https://www.paypalobjects.com/ppdevdocs/v1/img/docs/checkout/vertical_web.png"
        style="width: 100%; cursor: pointer;"
        @click="doAf()"
      >


    </div>

    <div class="mt-3" v-if="af">
      <b>April Fools!</b><br><br>

      Spire will always be free and open-source.
      <br><br>
      <b>Happy Spiring!</b><br><br>

      <button class="btn btn-primary" @click="closeAf()"><i class="fa fa-close"></i> Dismiss</button>
    </div>

    <template #modal-footer>
      <div class="">

      </div>
    </template>
  </b-modal>
</template>

<script>
import {LocalSettings, Setting} from "@/app/local-settings/localsettings";

export default {
  name: "Af-modal",
  data() {
    return {
      af: false,
    }
  },
  mounted() {
    // LocalSettings.set(Setting.AF, "false")
  },
  methods: {
    doAf() {
      this.af = true
      LocalSettings.set(Setting.AF, "true")
    },
    closeAf() {
      this.$bvModal.hide('af')
      // open rickroll
      window.open("https://www.youtube.com/watch?v=dQw4w9WgXcQ", "_blank")
    }
  }
}
</script>

<style scoped>

</style>
