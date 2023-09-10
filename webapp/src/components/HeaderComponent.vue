<template>
    <div class="header">
    <div class="speeddial-tooltip-demo" :style="{ position: 'relative', height: '350px' }">
        <SpeedDial :model="items" :radius="75" type="quarter-circle" showIcon="pi pi-bars" hideIcon="pi pi-times" direction="down-right"/>
    </div>
    </div>
</template>
<script>
import SpeedDial from 'primevue/speeddial';


export default {
    name: "HeaderComponent",
    components: {
        SpeedDial
    },
    data(){
        return {
            items: [
    {
        label: 'Add',
        icon: 'pi pi-plus',
        command: () => {
          this.$router.push({
            name: "noteCard", params: { opt: {
                Name: "",
                Description: "",
                Status: 1,
                isNew: true
              }
            }
          });
        }
    },
    {
        label: 'Delete',
        icon: 'pi pi-sign-out',
        command: () => {
            this.$router.push("/sign-in");
            localStorage.clear()
        }
    }
]
        }
    }, 
    created() {
        this.width = window.innerWidth;
        window.addEventListener('resize', this.updateWidth);
    },
    methods: {
        updateWidth() {
            this.width = window.innerWidth;
        }
    }
}
</script>
<style scoped>
.header {
    position: absolute;
    display: flex;
    flex-direction: row;
    align-content: space-between;
    background-color: #0012b4;
    margin: 0;
    padding: 5px;
    width: 100%;
    height: 50px;
    top:0;
}
.header-btns {
    display: flex;
    flex-direction: row;
    align-items: center;
    margin-left: auto;
}
.header-btn {
    margin-right: 1vw;
    border: 0;
}
.header-logo {
    height: 4vw;
    min-height: 45px;
}
.mobile{
    display: flex;
    position: fixed;
    bottom: 0;
    justify-content: space-between;
    z-index: 100;
    padding: 0;
    margin: 0;
    align-content: center;
    width: 100%;
    left: 0;
    background-color: #363537;
    box-shadow: 0 -5px 10px black;
}
</style>