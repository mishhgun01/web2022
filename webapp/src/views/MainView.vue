<template>
    <div class="container">
        <div v-for="(opt, idx) of options" :key="idx">
            <CardComponent  v-if="!selectedCard" :opt="opt" @openCard="openCard"/>
        </div>
    </div>
</template>
<script>
import CardComponent from "../components/CardComponent"
import {getNotes} from "@/api/notes";

export default {
    components: {
        CardComponent
    },
    data() {
        return {
            checked: null,
            options: [],
          selectedCard: null
        }
    },
    created() {
      const user = localStorage.getItem("User")
      if (!user) {
        this.$router.push("/sign-in")
        return
      }
      this.getData()
    },
  methods: {
      getData() {
        getNotes()
            .then(r => {
              this.options = r.data
            })
      },
    openCard(opt) {
        opt.isNew = false
        this.$router.push({name: "noteCard", params: { opt: opt } })
    }
  }
}
</script>
<style scoped>

.container {
    display: flex;
    justify-content:space-around;
    flex-direction: row;
    flex-wrap: wrap;
}
</style>