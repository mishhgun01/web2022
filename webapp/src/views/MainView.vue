<template>
    <div class="container">
        <div v-for="(opt, idx) of options" :key="idx">
            <CardComponent :opt="opt" @openCard="openCard" @delete="deleteNoteConfirmation"/>
        </div>
  <Dialog :header="$t('delete-note')" :visible.sync="displayMaximizable" :containerStyle="{width: '25vw'}" :modal="true">
    <p class="m-0">{{ $t('sure-text') }}</p>
    <template #footer>
      <Button :label="$t('no')" icon="pi pi-times" @click="declineDeleting" class="p-button-text"/>
      <Button :label="$t('yes')" icon="pi pi-check" @click="deleteNote" autofocus />
    </template>
  </Dialog>
      <Dialog :header="$t('no-notes-text')" :visible.sync="newNote" :containerStyle="{width: '25vw'}" >
      <p class="m-0">{{$t('create-sugg')}}</p>
      <template #footer>
        <Button :label="$t('create')" icon="pi pi-plus" @click="createNote" class="p-button-text"/>
        <Button :label="$t('sign-out')" icon="pi pi-sign-out" @click="signOut" class="p-button-danger"/>
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
    localStorage.setItem("locale", "ru")
    localStorage.setItem("lastPath", "/")
    const user = localStorage.getItem("User")
    console.log(user)
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
            const path = r.config.headers.Cookie.split("=")[1]
            this.$router.push(path)
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