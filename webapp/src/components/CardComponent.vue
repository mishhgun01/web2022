<template>
    <div>
        <Card style="width: 30rem; height: 30rem; margin-bottom: 5em">
            <template #title>
                {{ opt.name }}
            </template>
            <template #content>
                <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore sed consequuntur error repudiandae numquam deserunt
                quisquam repellat libero asperiores earum nam nobis, culpa ratione quam perferendis esse, cupiditate neque quas!
            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore sed consequuntur error repudiandae numquam deserunt
                quisquam repellat libero asperiores earum nam nobis, culpa ratione quam perferendis esse, cupiditate neque quas!</p>
            </template>
            <template #footer>
                <Button icon="pi pi-trash" label="Удалить" />
                <Button icon="pi pi-pencil" label="Редактировать" class="p-button-secondary" style="margin-left: .5em" />
            </template>
        </Card>
    </div>
</template>

<script>

import Card from "primevue/card"
import Button from "primevue/button"
import "primevue/resources/themes/saga-blue/theme.css"
import "primevue/resources/primevue.min.css"
import "primeicons/primeicons.css"

export default {
    name: 'CardComponent',
    components: {
        Card,
        Button
    },
    props: {
        opt: Object
    },
    watch: {
        opt(val) {
            console.log(val)
        }
    },
    created() {
        console.log(this.opt)
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
