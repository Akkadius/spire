<template>
  <div>
    <eq-window-simple title="Pet Selector">
      <b-input
        v-model="petSearch"
        class="form-control"
        v-on:keyup="searchPet"
        placeholder="Search by pet name..."
      />
    </eq-window-simple>

    <eq-window-simple
      id="pet-view-container"
      style="height: 85vh; overflow-y: scroll;" class="p-0"
    >
      <table
        id="pettable"
        class="eq-table eq-highlight-rows"
        style="display: table; font-size: 14px; overflow-x: scroll"
      >
        <thead
          class="eq-table-floating-header"
        >
        <tr>
          <th style="width: 30px"></th>
          <th style="width: 50px">NPC ID</th>
          <th>Name</th>
        </tr>
        </thead>
        <tbody>
        <tr
          :id="'pet-' + stripNonAlpha(pet.type)"
          :class="(isPetSelected(pet) ? 'pulsate-highlight-white' : '')"
          v-for="(pet, index) in filteredPets"
          :key="pet.id"
        >
          <td>
            <b-button
              class="btn-dark btn-sm btn-dark"
              @click="selectPet(pet)"
            >
              <i class="fa fa-arrow-left"></i>
            </b-button>
          </td>
          <td style="text-align: center">{{ pet.npc_id }}</td>
          <td style="text-align: left">{{ pet.type }}</td>
        </tr>
        </tbody>
      </table>
    </eq-window-simple>
  </div>
</template>

<script>
import {TELEPORT_ZONE_SELECTOR_TYPE} from "@/app/constants/eq-spell-constants";
import EqWindowSimple                from "@/components/eq-ui/EQWindowSimple";
import {PetApi}                      from "@/app/api";
import {SpireApi}              from "@/app/api/spire-api";
import util                          from "util";
import Expansions                    from "@/app/utility/expansions";
import EqCheckbox                    from "@/components/eq-ui/EQCheckbox";
import {SpireQueryBuilder}           from "@/app/api/spire-query-builder";

let pets = {}

export default {
  name: "SpellTeleportZoneSelectorPet",
  components: { EqCheckbox, EqWindowSimple },
  data() {
    return {
      TELEPORT_ZONE_SELECTOR_TYPE: TELEPORT_ZONE_SELECTOR_TYPE,

      // filtered content
      filteredPets: {},

      // search
      petSearch: "",

      // model we work with after the prop is passed so we can manipulate it ourselves
      // props should not be mutated
      selectedPet: "",
    }
  },
  props: {
    selectedPetName: {
      type: String,
      required: true,
    },
  },
  methods: {

    stripNonAlpha(string) {
      return string.replace(/[\W_]+/g, " ");
    },

    isPetSelected(pet) {
      return pet.type.trim() === this.selectedPet
    },

    selectPet(pet) {
      this.$emit('input', {
        pet: pet,
      });

      this.selectedPet = pet.type
    },

    searchPet() {
      const searchString = this.petSearch.toLowerCase().trim()
      let filteredPets   = []
      pets.forEach((pet) => {
        if (
          this.petSearch.trim() !== '' &&
          (
            pet.type.toLowerCase().includes(searchString)
          )) {
          filteredPets.push(pet)
        }
      });

      this.filteredPets = filteredPets

      if (filteredPets.length === 0) {
        this.filteredPets = pets;
      }
    },

    async loadPets() {
      const api    = (new PetApi(...SpireApi.cfg()))
      const result = await api.listPets(
        (new SpireQueryBuilder())
          .groupBy(["type"])
          .get()
      )

      if (result.status === 200) {
        pets              = result.data
        this.filteredPets = pets
      }
    },

    init() {
      this.loadPets()
    }
  },
  mounted() {
    // model we work with after the prop is passed - we can manipulate it ourselves
    this.selectedPet = this.selectedPetName
    this.init()

    setTimeout(() => {
      const container = document.getElementById("pet-view-container");
      const target    = document.getElementById(util.format("pet-%s", this.selectedPet))
      if (container && target) {
        const top           = target.getBoundingClientRect().top
        container.scrollTop = container.scrollTop + top - 300;
      }
    }, 1000)
  }
}
</script>

<style scoped>
#pettable td {
  vertical-align: middle !important;
}
</style>
