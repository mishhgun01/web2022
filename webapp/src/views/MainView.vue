<template>
    <div class="container">
        <div v-for="(opt, idx) of options" :key="idx">
            <CardComponent :opt="opt"/>
        </div>
    </div>
</template>
<script>
import CardComponent from "../components/CardComponent"

export default {
    components: {
        CardComponent
    },
    data() {
        return {
            checked: null,
            options: []
        }
    },
    created() {
        this.options.push(...[{name: "1", value: 1}, {name: "2", value: 2}, {name: "3", value: 3}, {name: "3", value: 3}, {name: "3", value: 3}, {name: "3", value: 3}])
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