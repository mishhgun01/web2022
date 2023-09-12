<template>
  <div>
    <Toast  position="top-right"/>
    <Card>
      <template #title>
        <div class="input w-50">
          <InputText v-model="name" class="p-inputtext-filled" type="text" placeholder="Название" :class="danger ? 'p-invalid' : ''"/>
        </div>
      </template>
      <template #content>
        <div class="input mt-5">
          <h3>Описание:</h3>
        </div>
        <Editor v-model="description" editorStyle="height: 320px" :class="danger ? 'p-invalid' : ''"/>
        <div class="input mt-5">
          <h3>Статус: {{ statusText }}</h3>
        </div>
        <div class="input mt-5">
          <Slider v-model="status" style="width: 14rem" :max="3" :min="1"/>
        </div>
      </template>
      <template #footer>
        <div class="input mt-5">
          <Button label="Сохранить" icon="pi pi-save" iconPos="right" class="p-button-raised p-button-rounded p-button-success" @click="saveNote"/>
        </div>
      </template>
    </Card>
  </div>
</template>

<script>
import Card from 'primevue/card'
import InputText from 'primevue/inputtext';
import Editor from 'primevue/editor';
import Slider from 'primevue/slider'
import Button from 'primevue/button'
import Toast from 'primevue/toast';
import {mappedStatus} from "@/helpers";
import {newNote, patchNote} from "@/api/notes";
import {updateUser} from "@/api/user";

export default {
  name: "NoteCard",
  components: {
    Card,
    InputText,
    Editor,
    Slider,
    Button,
    Toast
  },
  data() {
    return {
      ID: 0,
      danger: false,
      name: "",
      description: "",
      status: 1,
      userID: 0,
      isNew: false
    }
  },
  computed: {
    statusText() {
      return mappedStatus.find(s => s.status === this.status).value
    },
  },
  created() {
    localStorage.setItem("lastPath", this.$route.path)
    updateUser()
    const obj = this.$route.params.opt
    this.name = obj.Name || ""
    this.description = obj.Description || ""
    this.status = obj.Status || 1
    this.isNew = obj.isNew || ""
    this.ID = obj.ID || 0
  },
  methods: {
    saveNote() {
      if (!this.name?.length || !this.description?.length) {
        this.danger = true
        this.$toast.add({
          severity: 'error',
          summary: 'Ошибка',
          detail: 'Необходимо заполнить название и описание.',
          life: 3000
        });
      }
      if (this.isNew) {
        newNote({Name: this.name, Description: this.description, Status: this.status})
            .then(() => {
              this.$router.push("/")
            })
            .catch(() => {
              this.$toast.add({
                severity: 'error',
                summary: 'Ошибка',
                detail: 'Сервер не отвечает. Попробуйте позже.',
                life: 3000
              });
            })
        return
      } else {
        patchNote({ID: this.ID, Name: this.name, Description: this.description, Status: this.status}).then(() => {
          this.$router.push("/")
        })
            .catch(() => {
              this.$toast.add({
                severity: 'error',
                summary: 'Ошибка',
                detail: 'Сервер не отвечает. Попробуйте позже.',
                life: 3000
              });
            })
      }
    }
  }
}
</script>

<style scoped>

.input {
  display: flex;
  justify-content: start;
  align-items: center;
  margin-top: 5px;
  margin-right: 5px;
}
</style>