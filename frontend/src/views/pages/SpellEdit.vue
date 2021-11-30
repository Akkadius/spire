<template>
  <div class="container-fluid">
    <div class="panel-body">
      <div class="panel panel-default">
        <div class="row">
          <div class="col-7">
            <eq-window style="margin-top: 30px" title="Edit Spell">

              <eq-tabs v-if="spell">
                <eq-tab name="Basic" :selected="true">

                  Id
                  <b-form-input v-model="spell.id"/>
                  Name
                  <b-form-input v-model="spell.name"/>
                  You Cast
                  <b-form-input v-model="spell.you_cast"/>
                  Other Casts
                  <b-form-input v-model="spell.other_casts"/>
                  Cast On You
                  <b-form-input v-model="spell.cast_on_you"/>
                  Cast On Other
                  <b-form-input v-model="spell.cast_on_other"/>
                  Spell Fades
                  <b-form-input v-model="spell.spell_fades"/>
                  ID File
                  <b-form-input v-model="spell.player_1"/>


                </eq-tab>
                <eq-tab name="Effects">

                  <div >

                    <b-input-group style="height: 30px; margin-bottom: 15px">
                      <template #prepend>
                        <b-input-group-text style="width: 40px;">#</b-input-group-text>
                      </template>

                      <b-form-input placeholder="Effect" disabled style="width: 150px"/>
                      <b-form-input placeholder="Base" disabled/>
                      <b-form-input placeholder="Limit" disabled/>
                      <b-form-input placeholder="Max" disabled/>
                      <b-form-input placeholder="Formula" disabled/>
                    </b-input-group>

                    <b-input-group v-for="i in 12">
                      <template #prepend>
                        <b-input-group-text style="width: 40px;">{{ i }}</b-input-group-text>
                      </template>

                      <b-form-select v-model="spell['effectid_' + i]" @change="redrawCard" style="width: 150px">
                        <b-form-select-option v-for="(effect, id) in DB_SPA" :value="parseInt(id)">{{ id }}) {{
                            effect
                          }}
                        </b-form-select-option>
                      </b-form-select>

                      <b-form-input v-model="spell['effect_base_value_' + i]" @change="redrawCard"/>
                      <b-form-input v-model="spell['max_' + i]" @change="redrawCard"/>
                      <b-form-input v-model="spell['effect_limit_value_' + i]" @change="redrawCard"/>
                      <b-form-input v-model="spell['formula_' + i]" @change="redrawCard"/>
                    </b-input-group>

                    <!--                    {{i}})-->


                  </div>

                </eq-tab>
                <eq-tab name="Restrictions">

                </eq-tab>
              </eq-tabs>

            </eq-window>
          </div>

          <div class="col-5">
            <eq-window style="margin-top: 30px; margin-right: 10px; width: auto;" v-if="loaded">
              <eq-spell-preview :spell-data="spell"/>
            </eq-window>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import EqWindowFancy              from "../../components/eq-ui/EQWindowFancy";
import EqWindow                   from "../../components/eq-ui/EQWindow";
import EqTabs                     from "../../components/eq-ui/EQTabs";
import EqTab                      from "../../components/eq-ui/EQTab";
import EqSpellPreview             from "../../components/eq-ui/EQSpellCardPreview";
import {Spells}                   from "../../app/spells";
import {DB_SPA, DB_SPELL_EFFECTS} from "../../app/constants/eq-spell-constants";

export default {
  name: "SpellEdit",
  components: { EqSpellPreview, EqTab, EqTabs, EqWindow, EqWindowFancy },
  data() {
    return {
      spell: null,
      DB_SPELL_EFFECTS: DB_SPELL_EFFECTS,
      DB_SPA: DB_SPA,
      loaded: true,
    }
  },
  watch: {
    '$route'() {
      this.load()
    }
  },
  async created() {
    this.load()
    console.log("created")
  },
  methods: {
    load() {
      if (this.$route.params.id > 0) {
        Spells.getSpell(this.$route.params.id).then(result => {
          console.log(result)
          this.spell = result
        })
      }
    },
    redrawCard() {
      this.loaded = false

      setTimeout(() => {
        this.loaded = true
      }, 5)
    }
  }
}
</script>

<style scoped>

</style>
