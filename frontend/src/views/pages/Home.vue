<template>
  <div>

  </div>
</template>

<script>

  import EqWindow from "@/components/eq-ui/EQWindow";
  import UserContext from "@/app/user/UserContext";

  export default {
    components: { EqWindow },
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

