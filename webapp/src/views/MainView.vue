<template>
    <div class="container">
        <div v-for="(opt, idx) of options" :key="idx">
            <CardComponent :opt="opt" @openCard="openCard" @delete="deleteNoteConfirmation"/>
        </div>
  <Dialog header="Удаление заметки" :visible.sync="displayMaximizable" :containerStyle="{width: '25vw'}" :modal="true">
    <p class="m-0">Вы уверены?</p>
    <template #footer>
      <Button label="Нет" icon="pi pi-times" @click="declineDeleting" class="p-button-text"/>
      <Button label="Да" icon="pi pi-check" @click="deleteNote" autofocus />
    </template>
  </Dialog>
      <Dialog header="Ни одной заметки!" :visible.sync="newNote" :containerStyle="{width: '25vw'}" >
      <p class="m-0">Создайте новую заметку</p>
      <template #footer>
        <Button label="Создать" icon="pi pi-plus" @click="createNote" class="p-button-text"/>
        <Button label="Выйти" icon="pi pi-sign-out" @click="signOut" class="p-button-danger"/>
      </template>
    </Dialog>
  </div>
</template>
<script>
import CardComponent from "../components/CardComponent"
import Dialog from 'primevue/dialog'
import Button from "primevue/button";
export default {
  components: {
    CardComponent,
    Dialog,
    Button
  },
  data() {
    return {
      checked: null,
      options: [],
      selectedCard: null,
      displayMaximizable: false,
      newNote: false
    }
  },
  created() {
    document.cookie = "lastPath=/"
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
            if (!this.options?.length) {
              this.newNote = true
            }
          })
    },
    openCard(opt) {
      opt.isNew = false
      this.$router.push({name: "noteCard", params: { opt: opt } })
    },
    deleteNoteConfirmation(opt) {
      this.selectedCard = opt
      this.displayMaximizable = true

    },
    declineDeleting() {
      this.selectedCard = null
      this.displayMaximizable = false
    },
    deleteNote(){
      deleteNote(this.selectedCard)
          .then(()=> {
            this.getData()
            console.log("Ok")
            this.displayMaximizable = false
          })
    },
    createNote() {
      this.$router.push({
        name: "noteCard", params: { opt: {
            Name: "",
            Description: "",
            Status: 1,
            isNew: true
          }
        }
      })
    },
    signOut() {
      localStorage.clear()
      this.$router.push("/sign-in")
    }
  }
}

import {getNotes, deleteNote} from "@/api/notes";
</script>
<style scoped>

.container {
    display: flex;
    justify-content:space-around;
    flex-direction: row;
    flex-wrap: wrap;
}
</style>