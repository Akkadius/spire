package models

import (
	"github.com/volatiletech/null/v8"
	"time"
)

type Item struct {
	ID                      int                     `json:"id" gorm:"Column:id"`
	Minstatus               int16                   `json:"minstatus" gorm:"Column:minstatus"`
	Name                    string                  `json:"name" gorm:"Column:Name"`
	Aagi                    int                     `json:"aagi" gorm:"Column:aagi"`
	Ac                      int                     `json:"ac" gorm:"Column:ac"`
	Accuracy                int                     `json:"accuracy" gorm:"Column:accuracy"`
	Acha                    int                     `json:"acha" gorm:"Column:acha"`
	Adex                    int                     `json:"adex" gorm:"Column:adex"`
	Aint                    int                     `json:"aint" gorm:"Column:aint"`
	Artifactflag            uint8                   `json:"artifactflag" gorm:"Column:artifactflag"`
	Asta                    int                     `json:"asta" gorm:"Column:asta"`
	Astr                    int                     `json:"astr" gorm:"Column:astr"`
	Attack                  int                     `json:"attack" gorm:"Column:attack"`
	Augrestrict             int                     `json:"augrestrict" gorm:"Column:augrestrict"`
	Augslot1Type            int8                    `json:"augslot_1_type" gorm:"Column:augslot1type"`
	Augslot1Visible         int8                    `json:"augslot_1_visible" gorm:"Column:augslot1visible"`
	Augslot2Type            int8                    `json:"augslot_2_type" gorm:"Column:augslot2type"`
	Augslot2Visible         int8                    `json:"augslot_2_visible" gorm:"Column:augslot2visible"`
	Augslot3Type            int8                    `json:"augslot_3_type" gorm:"Column:augslot3type"`
	Augslot3Visible         int8                    `json:"augslot_3_visible" gorm:"Column:augslot3visible"`
	Augslot4Type            int8                    `json:"augslot_4_type" gorm:"Column:augslot4type"`
	Augslot4Visible         int8                    `json:"augslot_4_visible" gorm:"Column:augslot4visible"`
	Augslot5Type            int8                    `json:"augslot_5_type" gorm:"Column:augslot5type"`
	Augslot5Visible         int8                    `json:"augslot_5_visible" gorm:"Column:augslot5visible"`
	Augslot6Type            int8                    `json:"augslot_6_type" gorm:"Column:augslot6type"`
	Augslot6Visible         int8                    `json:"augslot_6_visible" gorm:"Column:augslot6visible"`
	Augtype                 int                     `json:"augtype" gorm:"Column:augtype"`
	Avoidance               int                     `json:"avoidance" gorm:"Column:avoidance"`
	Awis                    int                     `json:"awis" gorm:"Column:awis"`
	Bagsize                 int                     `json:"bagsize" gorm:"Column:bagsize"`
	Bagslots                int                     `json:"bagslots" gorm:"Column:bagslots"`
	Bagtype                 int                     `json:"bagtype" gorm:"Column:bagtype"`
	Bagwr                   int                     `json:"bagwr" gorm:"Column:bagwr"`
	Banedmgamt              int                     `json:"banedmgamt" gorm:"Column:banedmgamt"`
	Banedmgraceamt          int                     `json:"banedmgraceamt" gorm:"Column:banedmgraceamt"`
	Banedmgbody             int                     `json:"banedmgbody" gorm:"Column:banedmgbody"`
	Banedmgrace             int                     `json:"banedmgrace" gorm:"Column:banedmgrace"`
	Bardtype                int                     `json:"bardtype" gorm:"Column:bardtype"`
	Bardvalue               int                     `json:"bardvalue" gorm:"Column:bardvalue"`
	Book                    int                     `json:"book" gorm:"Column:book"`
	Casttime                int                     `json:"casttime" gorm:"Column:casttime"`
	Casttime2               int                     `json:"casttime_" gorm:"Column:casttime_"`
	Charmfile               string                  `json:"charmfile" gorm:"Column:charmfile"`
	Charmfileid             string                  `json:"charmfileid" gorm:"Column:charmfileid"`
	Classes                 int                     `json:"classes" gorm:"Column:classes"`
	Color                   uint                    `json:"color" gorm:"Column:color"`
	Combateffects           string                  `json:"combateffects" gorm:"Column:combateffects"`
	Extradmgskill           int                     `json:"extradmgskill" gorm:"Column:extradmgskill"`
	Extradmgamt             int                     `json:"extradmgamt" gorm:"Column:extradmgamt"`
	Price                   int                     `json:"price" gorm:"Column:price"`
	Cr                      int                     `json:"cr" gorm:"Column:cr"`
	Damage                  int                     `json:"damage" gorm:"Column:damage"`
	Damageshield            int                     `json:"damageshield" gorm:"Column:damageshield"`
	Deity                   int                     `json:"deity" gorm:"Column:deity"`
	Delay                   int                     `json:"delay" gorm:"Column:delay"`
	Augdistiller            int                     `json:"augdistiller" gorm:"Column:augdistiller"`
	Dotshielding            int                     `json:"dotshielding" gorm:"Column:dotshielding"`
	Dr                      int                     `json:"dr" gorm:"Column:dr"`
	Clicktype               int                     `json:"clicktype" gorm:"Column:clicktype"`
	Clicklevel2             int                     `json:"clicklevel_2" gorm:"Column:clicklevel2"`
	Elemdmgtype             int                     `json:"elemdmgtype" gorm:"Column:elemdmgtype"`
	Elemdmgamt              int                     `json:"elemdmgamt" gorm:"Column:elemdmgamt"`
	Endur                   int                     `json:"endur" gorm:"Column:endur"`
	Factionamt1             int                     `json:"factionamt_1" gorm:"Column:factionamt1"`
	Factionamt2             int                     `json:"factionamt_2" gorm:"Column:factionamt2"`
	Factionamt3             int                     `json:"factionamt_3" gorm:"Column:factionamt3"`
	Factionamt4             int                     `json:"factionamt_4" gorm:"Column:factionamt4"`
	Factionmod1             int                     `json:"factionmod_1" gorm:"Column:factionmod1"`
	Factionmod2             int                     `json:"factionmod_2" gorm:"Column:factionmod2"`
	Factionmod3             int                     `json:"factionmod_3" gorm:"Column:factionmod3"`
	Factionmod4             int                     `json:"factionmod_4" gorm:"Column:factionmod4"`
	Filename                string                  `json:"filename" gorm:"Column:filename"`
	Focuseffect             int                     `json:"focuseffect" gorm:"Column:focuseffect"`
	Fr                      int                     `json:"fr" gorm:"Column:fr"`
	Fvnodrop                int                     `json:"fvnodrop" gorm:"Column:fvnodrop"`
	Haste                   int                     `json:"haste" gorm:"Column:haste"`
	Clicklevel              int                     `json:"clicklevel" gorm:"Column:clicklevel"`
	Hp                      int                     `json:"hp" gorm:"Column:hp"`
	Regen                   int                     `json:"regen" gorm:"Column:regen"`
	Icon                    int                     `json:"icon" gorm:"Column:icon"`
	Idfile                  string                  `json:"idfile" gorm:"Column:idfile"`
	Itemclass               int                     `json:"itemclass" gorm:"Column:itemclass"`
	Itemtype                int                     `json:"itemtype" gorm:"Column:itemtype"`
	Ldonprice               int                     `json:"ldonprice" gorm:"Column:ldonprice"`
	Ldontheme               int                     `json:"ldontheme" gorm:"Column:ldontheme"`
	Ldonsold                int                     `json:"ldonsold" gorm:"Column:ldonsold"`
	Light                   int                     `json:"light" gorm:"Column:light"`
	Lore                    string                  `json:"lore" gorm:"Column:lore"`
	Loregroup               int                     `json:"loregroup" gorm:"Column:loregroup"`
	Magic                   int                     `json:"magic" gorm:"Column:magic"`
	Mana                    int                     `json:"mana" gorm:"Column:mana"`
	Manaregen               int                     `json:"manaregen" gorm:"Column:manaregen"`
	Enduranceregen          int                     `json:"enduranceregen" gorm:"Column:enduranceregen"`
	Material                int                     `json:"material" gorm:"Column:material"`
	Herosforgemodel         int                     `json:"herosforgemodel" gorm:"Column:herosforgemodel"`
	Maxcharges              int                     `json:"maxcharges" gorm:"Column:maxcharges"`
	Mr                      int                     `json:"mr" gorm:"Column:mr"`
	Nodrop                  int                     `json:"nodrop" gorm:"Column:nodrop"`
	Norent                  int                     `json:"norent" gorm:"Column:norent"`
	Pendingloreflag         uint8                   `json:"pendingloreflag" gorm:"Column:pendingloreflag"`
	Pr                      int                     `json:"pr" gorm:"Column:pr"`
	Procrate                int                     `json:"procrate" gorm:"Column:procrate"`
	Races                   int                     `json:"races" gorm:"Column:races"`
	Range                   int                     `json:"range" gorm:"Column:range"`
	Reclevel                int                     `json:"reclevel" gorm:"Column:reclevel"`
	Recskill                int                     `json:"recskill" gorm:"Column:recskill"`
	Reqlevel                int                     `json:"reqlevel" gorm:"Column:reqlevel"`
	Sellrate                float32                 `json:"sellrate" gorm:"Column:sellrate"`
	Shielding               int                     `json:"shielding" gorm:"Column:shielding"`
	Size                    int                     `json:"size" gorm:"Column:size"`
	Skillmodtype            int                     `json:"skillmodtype" gorm:"Column:skillmodtype"`
	Skillmodvalue           int                     `json:"skillmodvalue" gorm:"Column:skillmodvalue"`
	Slots                   int                     `json:"slots" gorm:"Column:slots"`
	Clickeffect             int                     `json:"clickeffect" gorm:"Column:clickeffect"`
	Spellshield             int                     `json:"spellshield" gorm:"Column:spellshield"`
	Strikethrough           int                     `json:"strikethrough" gorm:"Column:strikethrough"`
	Stunresist              int                     `json:"stunresist" gorm:"Column:stunresist"`
	Summonedflag            uint8                   `json:"summonedflag" gorm:"Column:summonedflag"`
	Tradeskills             int                     `json:"tradeskills" gorm:"Column:tradeskills"`
	Favor                   int                     `json:"favor" gorm:"Column:favor"`
	Weight                  int                     `json:"weight" gorm:"Column:weight"`
	UNK012                  int                     `json:"unk_012" gorm:"Column:UNK012"`
	UNK013                  int                     `json:"unk_013" gorm:"Column:UNK013"`
	Benefitflag             int                     `json:"benefitflag" gorm:"Column:benefitflag"`
	UNK054                  int                     `json:"unk_054" gorm:"Column:UNK054"`
	UNK059                  int                     `json:"unk_059" gorm:"Column:UNK059"`
	Booktype                int                     `json:"booktype" gorm:"Column:booktype"`
	Recastdelay             int                     `json:"recastdelay" gorm:"Column:recastdelay"`
	Recasttype              int                     `json:"recasttype" gorm:"Column:recasttype"`
	Guildfavor              int                     `json:"guildfavor" gorm:"Column:guildfavor"`
	UNK123                  int                     `json:"unk_123" gorm:"Column:UNK123"`
	UNK124                  int                     `json:"unk_124" gorm:"Column:UNK124"`
	Attuneable              int                     `json:"attuneable" gorm:"Column:attuneable"`
	Nopet                   int                     `json:"nopet" gorm:"Column:nopet"`
	Updated                 time.Time               `json:"updated" gorm:"Column:updated"`
	Comment                 string                  `json:"comment" gorm:"Column:comment"`
	UNK127                  int                     `json:"unk_127" gorm:"Column:UNK127"`
	Pointtype               int                     `json:"pointtype" gorm:"Column:pointtype"`
	Potionbelt              int                     `json:"potionbelt" gorm:"Column:potionbelt"`
	Potionbeltslots         int                     `json:"potionbeltslots" gorm:"Column:potionbeltslots"`
	Stacksize               int                     `json:"stacksize" gorm:"Column:stacksize"`
	Notransfer              int                     `json:"notransfer" gorm:"Column:notransfer"`
	Stackable               int                     `json:"stackable" gorm:"Column:stackable"`
	UNK134                  string                  `json:"unk_134" gorm:"Column:UNK134"`
	UNK137                  int                     `json:"unk_137" gorm:"Column:UNK137"`
	Proceffect              int                     `json:"proceffect" gorm:"Column:proceffect"`
	Proctype                int                     `json:"proctype" gorm:"Column:proctype"`
	Proclevel2              int                     `json:"proclevel_2" gorm:"Column:proclevel2"`
	Proclevel               int                     `json:"proclevel" gorm:"Column:proclevel"`
	UNK142                  int                     `json:"unk_142" gorm:"Column:UNK142"`
	Worneffect              int                     `json:"worneffect" gorm:"Column:worneffect"`
	Worntype                int                     `json:"worntype" gorm:"Column:worntype"`
	Wornlevel2              int                     `json:"wornlevel_2" gorm:"Column:wornlevel2"`
	Wornlevel               int                     `json:"wornlevel" gorm:"Column:wornlevel"`
	UNK147                  int                     `json:"unk_147" gorm:"Column:UNK147"`
	Focustype               int                     `json:"focustype" gorm:"Column:focustype"`
	Focuslevel2             int                     `json:"focuslevel_2" gorm:"Column:focuslevel2"`
	Focuslevel              int                     `json:"focuslevel" gorm:"Column:focuslevel"`
	UNK152                  int                     `json:"unk_152" gorm:"Column:UNK152"`
	Scrolleffect            int                     `json:"scrolleffect" gorm:"Column:scrolleffect"`
	Scrolltype              int                     `json:"scrolltype" gorm:"Column:scrolltype"`
	Scrolllevel2            int                     `json:"scrolllevel_2" gorm:"Column:scrolllevel2"`
	Scrolllevel             int                     `json:"scrolllevel" gorm:"Column:scrolllevel"`
	UNK157                  int                     `json:"unk_157" gorm:"Column:UNK157"`
	Serialized              null.Time               `json:"serialized" gorm:"Column:serialized"`
	Verified                null.Time               `json:"verified" gorm:"Column:verified"`
	Serialization           null.String             `json:"serialization" gorm:"Column:serialization"`
	Source                  string                  `json:"source" gorm:"Column:source"`
	UNK033                  int                     `json:"unk_033" gorm:"Column:UNK033"`
	Lorefile                string                  `json:"lorefile" gorm:"Column:lorefile"`
	UNK014                  int                     `json:"unk_014" gorm:"Column:UNK014"`
	Svcorruption            int                     `json:"svcorruption" gorm:"Column:svcorruption"`
	Skillmodmax             int                     `json:"skillmodmax" gorm:"Column:skillmodmax"`
	UNK060                  int                     `json:"unk_060" gorm:"Column:UNK060"`
	Augslot1Unk2            int                     `json:"augslot_1_unk_2" gorm:"Column:augslot1unk2"`
	Augslot2Unk2            int                     `json:"augslot_2_unk_2" gorm:"Column:augslot2unk2"`
	Augslot3Unk2            int                     `json:"augslot_3_unk_2" gorm:"Column:augslot3unk2"`
	Augslot4Unk2            int                     `json:"augslot_4_unk_2" gorm:"Column:augslot4unk2"`
	Augslot5Unk2            int                     `json:"augslot_5_unk_2" gorm:"Column:augslot5unk2"`
	Augslot6Unk2            int                     `json:"augslot_6_unk_2" gorm:"Column:augslot6unk2"`
	UNK120                  int                     `json:"unk_120" gorm:"Column:UNK120"`
	UNK121                  int                     `json:"unk_121" gorm:"Column:UNK121"`
	Questitemflag           int                     `json:"questitemflag" gorm:"Column:questitemflag"`
	UNK132                  null.String             `json:"unk_132" gorm:"Column:UNK132"`
	Clickunk5               int                     `json:"clickunk_5" gorm:"Column:clickunk5"`
	Clickunk6               string                  `json:"clickunk_6" gorm:"Column:clickunk6"`
	Clickunk7               int                     `json:"clickunk_7" gorm:"Column:clickunk7"`
	Procunk1                int                     `json:"procunk_1" gorm:"Column:procunk1"`
	Procunk2                int                     `json:"procunk_2" gorm:"Column:procunk2"`
	Procunk3                int                     `json:"procunk_3" gorm:"Column:procunk3"`
	Procunk4                int                     `json:"procunk_4" gorm:"Column:procunk4"`
	Procunk6                string                  `json:"procunk_6" gorm:"Column:procunk6"`
	Procunk7                int                     `json:"procunk_7" gorm:"Column:procunk7"`
	Wornunk1                int                     `json:"wornunk_1" gorm:"Column:wornunk1"`
	Wornunk2                int                     `json:"wornunk_2" gorm:"Column:wornunk2"`
	Wornunk3                int                     `json:"wornunk_3" gorm:"Column:wornunk3"`
	Wornunk4                int                     `json:"wornunk_4" gorm:"Column:wornunk4"`
	Wornunk5                int                     `json:"wornunk_5" gorm:"Column:wornunk5"`
	Wornunk6                string                  `json:"wornunk_6" gorm:"Column:wornunk6"`
	Wornunk7                int                     `json:"wornunk_7" gorm:"Column:wornunk7"`
	Focusunk1               int                     `json:"focusunk_1" gorm:"Column:focusunk1"`
	Focusunk2               int                     `json:"focusunk_2" gorm:"Column:focusunk2"`
	Focusunk3               int                     `json:"focusunk_3" gorm:"Column:focusunk3"`
	Focusunk4               int                     `json:"focusunk_4" gorm:"Column:focusunk4"`
	Focusunk5               int                     `json:"focusunk_5" gorm:"Column:focusunk5"`
	Focusunk6               string                  `json:"focusunk_6" gorm:"Column:focusunk6"`
	Focusunk7               int                     `json:"focusunk_7" gorm:"Column:focusunk7"`
	Scrollunk1              int                     `json:"scrollunk_1" gorm:"Column:scrollunk1"`
	Scrollunk2              int                     `json:"scrollunk_2" gorm:"Column:scrollunk2"`
	Scrollunk3              int                     `json:"scrollunk_3" gorm:"Column:scrollunk3"`
	Scrollunk4              int                     `json:"scrollunk_4" gorm:"Column:scrollunk4"`
	Scrollunk5              int                     `json:"scrollunk_5" gorm:"Column:scrollunk5"`
	Scrollunk6              string                  `json:"scrollunk_6" gorm:"Column:scrollunk6"`
	Scrollunk7              int                     `json:"scrollunk_7" gorm:"Column:scrollunk7"`
	UNK193                  int                     `json:"unk_193" gorm:"Column:UNK193"`
	Purity                  int                     `json:"purity" gorm:"Column:purity"`
	Evoitem                 int                     `json:"evoitem" gorm:"Column:evoitem"`
	Evoid                   int                     `json:"evoid" gorm:"Column:evoid"`
	Evolvinglevel           int                     `json:"evolvinglevel" gorm:"Column:evolvinglevel"`
	Evomax                  int                     `json:"evomax" gorm:"Column:evomax"`
	Clickname               string                  `json:"clickname" gorm:"Column:clickname"`
	Procname                string                  `json:"procname" gorm:"Column:procname"`
	Wornname                string                  `json:"wornname" gorm:"Column:wornname"`
	Focusname               string                  `json:"focusname" gorm:"Column:focusname"`
	Scrollname              string                  `json:"scrollname" gorm:"Column:scrollname"`
	Dsmitigation            int16                   `json:"dsmitigation" gorm:"Column:dsmitigation"`
	HeroicStr               int16                   `json:"heroic_str" gorm:"Column:heroic_str"`
	HeroicInt               int16                   `json:"heroic_int" gorm:"Column:heroic_int"`
	HeroicWis               int16                   `json:"heroic_wis" gorm:"Column:heroic_wis"`
	HeroicAgi               int16                   `json:"heroic_agi" gorm:"Column:heroic_agi"`
	HeroicDex               int16                   `json:"heroic_dex" gorm:"Column:heroic_dex"`
	HeroicSta               int16                   `json:"heroic_sta" gorm:"Column:heroic_sta"`
	HeroicCha               int16                   `json:"heroic_cha" gorm:"Column:heroic_cha"`
	HeroicPr                int16                   `json:"heroic_pr" gorm:"Column:heroic_pr"`
	HeroicDr                int16                   `json:"heroic_dr" gorm:"Column:heroic_dr"`
	HeroicFr                int16                   `json:"heroic_fr" gorm:"Column:heroic_fr"`
	HeroicCr                int16                   `json:"heroic_cr" gorm:"Column:heroic_cr"`
	HeroicMr                int16                   `json:"heroic_mr" gorm:"Column:heroic_mr"`
	HeroicSvcorrup          int16                   `json:"heroic_svcorrup" gorm:"Column:heroic_svcorrup"`
	Healamt                 int16                   `json:"healamt" gorm:"Column:healamt"`
	Spelldmg                int16                   `json:"spelldmg" gorm:"Column:spelldmg"`
	Clairvoyance            int16                   `json:"clairvoyance" gorm:"Column:clairvoyance"`
	Backstabdmg             int16                   `json:"backstabdmg" gorm:"Column:backstabdmg"`
	Created                 string                  `json:"created" gorm:"Column:created"`
	Elitematerial           int16                   `json:"elitematerial" gorm:"Column:elitematerial"`
	Ldonsellbackrate        int16                   `json:"ldonsellbackrate" gorm:"Column:ldonsellbackrate"`
	Scriptfileid            int16                   `json:"scriptfileid" gorm:"Column:scriptfileid"`
	Expendablearrow         int16                   `json:"expendablearrow" gorm:"Column:expendablearrow"`
	Powersourcecapacity     int16                   `json:"powersourcecapacity" gorm:"Column:powersourcecapacity"`
	Bardeffect              int16                   `json:"bardeffect" gorm:"Column:bardeffect"`
	Bardeffecttype          int16                   `json:"bardeffecttype" gorm:"Column:bardeffecttype"`
	Bardlevel2              int16                   `json:"bardlevel_2" gorm:"Column:bardlevel2"`
	Bardlevel               int16                   `json:"bardlevel" gorm:"Column:bardlevel"`
	Bardunk1                int16                   `json:"bardunk_1" gorm:"Column:bardunk1"`
	Bardunk2                int16                   `json:"bardunk_2" gorm:"Column:bardunk2"`
	Bardunk3                int16                   `json:"bardunk_3" gorm:"Column:bardunk3"`
	Bardunk4                int16                   `json:"bardunk_4" gorm:"Column:bardunk4"`
	Bardunk5                int16                   `json:"bardunk_5" gorm:"Column:bardunk5"`
	Bardname                string                  `json:"bardname" gorm:"Column:bardname"`
	Bardunk7                int16                   `json:"bardunk_7" gorm:"Column:bardunk7"`
	UNK214                  int16                   `json:"unk_214" gorm:"Column:UNK214"`
	Subtype                 int                     `json:"subtype" gorm:"Column:subtype"`
	UNK220                  int                     `json:"unk_220" gorm:"Column:UNK220"`
	UNK221                  int                     `json:"unk_221" gorm:"Column:UNK221"`
	Heirloom                int                     `json:"heirloom" gorm:"Column:heirloom"`
	UNK223                  int                     `json:"unk_223" gorm:"Column:UNK223"`
	UNK224                  int                     `json:"unk_224" gorm:"Column:UNK224"`
	UNK225                  int                     `json:"unk_225" gorm:"Column:UNK225"`
	UNK226                  int                     `json:"unk_226" gorm:"Column:UNK226"`
	UNK227                  int                     `json:"unk_227" gorm:"Column:UNK227"`
	UNK228                  int                     `json:"unk_228" gorm:"Column:UNK228"`
	UNK229                  int                     `json:"unk_229" gorm:"Column:UNK229"`
	UNK230                  int                     `json:"unk_230" gorm:"Column:UNK230"`
	UNK231                  int                     `json:"unk_231" gorm:"Column:UNK231"`
	UNK232                  int                     `json:"unk_232" gorm:"Column:UNK232"`
	UNK233                  int                     `json:"unk_233" gorm:"Column:UNK233"`
	UNK234                  int                     `json:"unk_234" gorm:"Column:UNK234"`
	Placeable               int                     `json:"placeable" gorm:"Column:placeable"`
	UNK236                  int                     `json:"unk_236" gorm:"Column:UNK236"`
	UNK237                  int                     `json:"unk_237" gorm:"Column:UNK237"`
	UNK238                  int                     `json:"unk_238" gorm:"Column:UNK238"`
	UNK239                  int                     `json:"unk_239" gorm:"Column:UNK239"`
	UNK240                  int                     `json:"unk_240" gorm:"Column:UNK240"`
	UNK241                  int                     `json:"unk_241" gorm:"Column:UNK241"`
	Epicitem                int                     `json:"epicitem" gorm:"Column:epicitem"`
	AlternateCurrencies     []AlternateCurrency     `json:"alternate_currencies,omitempty" gorm:"foreignKey:item_id;references:id"`
	CharacterCorpseItems    []CharacterCorpseItem   `json:"character_corpse_items,omitempty" gorm:"foreignKey:item_id;references:id"`
	DiscoveredItems         []DiscoveredItem        `json:"discovered_items,omitempty" gorm:"foreignKey:item_id;references:id"`
	Doors                   []Door                  `json:"doors,omitempty" gorm:"foreignKey:keyitem;references:id"`
	Fishings                []Fishing               `json:"fishings,omitempty" gorm:"foreignKey:Itemid;references:id"`
	Forages                 []Forage                `json:"forages,omitempty" gorm:"foreignKey:Itemid;references:id"`
	ItemTicks               []ItemTick              `json:"item_ticks,omitempty" gorm:"foreignKey:it_itemid;references:id"`
	Keyrings                []Keyring               `json:"keyrings,omitempty" gorm:"foreignKey:item_id;references:id"`
	LootdropEntries         []LootdropEntry         `json:"lootdrop_entries,omitempty" gorm:"foreignKey:item_id;references:id"`
	Objects                 []Object                `json:"objects,omitempty" gorm:"foreignKey:itemid;references:id"`
	ObjectContents          []ObjectContent         `json:"object_contents,omitempty" gorm:"foreignKey:itemid;references:id"`
	StartingItems           []StartingItem          `json:"starting_items,omitempty" gorm:"foreignKey:itemid;references:id"`
	TradeskillRecipeEntries []TradeskillRecipeEntry `json:"tradeskill_recipe_entries,omitempty" gorm:"foreignKey:item_id;references:id"`
	TributeLevels           []TributeLevel          `json:"tribute_levels,omitempty" gorm:"foreignKey:item_id;references:id"`
	GroundSpawns            []GroundSpawn           `json:"ground_spawns,omitempty" gorm:"foreignKey:item;references:id"`
	Merchantlists           []Merchantlist          `json:"merchantlists,omitempty" gorm:"foreignKey:item;references:id"`
	Tasks                   []Task                  `json:"tasks,omitempty" gorm:"foreignKey:rewardid;references:id"`
}

func (Item) TableName() string {
    return "items"
}

func (Item) Relationships() []string {
    return []string{
		"AlternateCurrencies",
		"AlternateCurrencies.Item",
		"CharacterCorpseItems",
		"DiscoveredItems",
		"Doors",
		"Doors.Item",
		"Fishings",
		"Fishings.Item",
		"Fishings.NpcType",
		"Fishings.NpcType.AlternateCurrency",
		"Fishings.NpcType.AlternateCurrency.Item",
		"Fishings.NpcType.Loottable",
		"Fishings.NpcType.Loottable.LoottableEntries",
		"Fishings.NpcType.Loottable.LoottableEntries.Lootdrop",
		"Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries",
		"Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item",
		"Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop",
		"Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries",
		"Fishings.NpcType.Loottable.LoottableEntries.Loottable",
		"Fishings.NpcType.Loottable.NpcTypes",
		"Fishings.NpcType.Merchantlists",
		"Fishings.NpcType.Merchantlists.Items",
		"Fishings.NpcType.Merchantlists.NpcTypes",
		"Fishings.NpcType.NpcEmotes",
		"Fishings.NpcType.NpcFactions",
		"Fishings.NpcType.NpcFactions.NpcFactionEntries",
		"Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList",
		"Fishings.NpcType.NpcSpell",
		"Fishings.NpcType.NpcSpell.NpcSpellsEntries",
		"Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew",
		"Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura",
		"Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items",
		"Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"Fishings.NpcType.NpcTypesTint",
		"Fishings.NpcType.Spawnentries",
		"Fishings.NpcType.Spawnentries.NpcType",
		"Fishings.NpcType.Spawnentries.Spawngroup",
		"Fishings.NpcType.Spawnentries.Spawngroup.Spawn2",
		"Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Fishings.Zone",
		"Forages",
		"Forages.Item",
		"Forages.Zone",
		"GroundSpawns",
		"GroundSpawns.Zone",
		"ItemTicks",
		"Keyrings",
		"LootdropEntries",
		"LootdropEntries.Item",
		"LootdropEntries.Lootdrop",
		"LootdropEntries.Lootdrop.LootdropEntries",
		"LootdropEntries.Lootdrop.LoottableEntries",
		"LootdropEntries.Lootdrop.LoottableEntries.Lootdrop",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Merchantlists",
		"Merchantlists.Items",
		"Merchantlists.NpcTypes",
		"Merchantlists.NpcTypes.AlternateCurrency",
		"Merchantlists.NpcTypes.AlternateCurrency.Item",
		"Merchantlists.NpcTypes.Loottable",
		"Merchantlists.NpcTypes.Loottable.LoottableEntries",
		"Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop",
		"Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries",
		"Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item",
		"Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop",
		"Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries",
		"Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable",
		"Merchantlists.NpcTypes.Loottable.NpcTypes",
		"Merchantlists.NpcTypes.Merchantlists",
		"Merchantlists.NpcTypes.NpcEmotes",
		"Merchantlists.NpcTypes.NpcFactions",
		"Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries",
		"Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList",
		"Merchantlists.NpcTypes.NpcSpell",
		"Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries",
		"Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew",
		"Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura",
		"Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items",
		"Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"Merchantlists.NpcTypes.NpcTypesTint",
		"Merchantlists.NpcTypes.Spawnentries",
		"Merchantlists.NpcTypes.Spawnentries.NpcType",
		"Merchantlists.NpcTypes.Spawnentries.Spawngroup",
		"Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"ObjectContents",
		"Objects",
		"Objects.Item",
		"Objects.Zone",
		"StartingItems",
		"StartingItems.Item",
		"StartingItems.Zone",
		"Tasks",
		"Tasks.AlternateCurrency",
		"Tasks.AlternateCurrency.Item",
		"Tasks.TaskActivities",
		"Tasks.TaskActivities.Goallists",
		"Tasks.TaskActivities.NpcType",
		"Tasks.TaskActivities.NpcType.AlternateCurrency",
		"Tasks.TaskActivities.NpcType.AlternateCurrency.Item",
		"Tasks.TaskActivities.NpcType.Loottable",
		"Tasks.TaskActivities.NpcType.Loottable.LoottableEntries",
		"Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop",
		"Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries",
		"Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item",
		"Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop",
		"Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries",
		"Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable",
		"Tasks.TaskActivities.NpcType.Loottable.NpcTypes",
		"Tasks.TaskActivities.NpcType.Merchantlists",
		"Tasks.TaskActivities.NpcType.Merchantlists.Items",
		"Tasks.TaskActivities.NpcType.Merchantlists.NpcTypes",
		"Tasks.TaskActivities.NpcType.NpcEmotes",
		"Tasks.TaskActivities.NpcType.NpcFactions",
		"Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries",
		"Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries.FactionList",
		"Tasks.TaskActivities.NpcType.NpcSpell",
		"Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries",
		"Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew",
		"Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura",
		"Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items",
		"Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"Tasks.TaskActivities.NpcType.NpcTypesTint",
		"Tasks.TaskActivities.NpcType.Spawnentries",
		"Tasks.TaskActivities.NpcType.Spawnentries.NpcType",
		"Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup",
		"Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2",
		"Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Tasks.Tasksets",
		"TradeskillRecipeEntries",
		"TradeskillRecipeEntries.TradeskillRecipe",
		"TributeLevels",
	}
}

func (Item) Connection() string {
    return "eqemu_content"
}
