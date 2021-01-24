<template>
  <div>

  </div>
</template>

<script>

  import HeaderImageComponent from '../../components/HeaderImageComponent';
  import EqWindow from "@/components/eq-ui/EQWindow";
  import UserContext from "@/app/user/UserContext";

  export default {
    components: { EqWindow, HeaderImageComponent },
    data() {
      return {
        userContext: null,
      }
    },
    async mounted() {
      this.userContext = await (UserContext.getUser())
    }


  }
</script>

