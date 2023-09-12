<template>
    <div>
        <Card style="width: 30rem; height: 30rem; margin-bottom: 5em">
            <template #title>
                {{ opt.Name }}
            </template>
            <template #content>
                <Editor :value="opt.Description" :readonly="true" editorStyle="height: 250px"/>
            </template>
            <template #footer>
                <Button icon="pi pi-trash" :label="$t('delete')" @click="deleteNote"/>
                <Button icon="pi pi-pencil" :label="$t('edit')" class="p-button-secondary" style="margin-left: .5em" @click="openCard"/>
            </template>
        </Card>
    </div>
</template>

<script>

import Card from "primevue/card"
import Button from "primevue/button"
import Editor from "primevue/editor";
import "primevue/resources/themes/saga-blue/theme.css"
import "primevue/resources/primevue.min.css"
import "primeicons/primeicons.css"

export default {
    name: 'CardComponent',
    components: {
        Card,
        Button,
      Editor
    },
    props: {
        opt: Object
    },
  methods: {
      openCard() {
        this.$emit("openCard", this.opt)
      },
    deleteNote() {
        this.$emit("delete", this.opt)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
margin: 40px 0 0;
}
ul {
list-style-type: none;
padding: 0;
}
li {
display: inline-block;
margin: 0 10px;
}
a {
color: #42b983;
}
</style>
