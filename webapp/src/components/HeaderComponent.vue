<template>
    <div class="header">
    <div class="speeddial-tooltip-demo" :style="{ position: 'relative', height: '400px' }">
        <SpeedDial :model="items" :radius="60" showIcon="pi pi-bars" hideIcon="pi pi-times" direction="down"/>
      <div class="title">
        <h3>Мои заметки</h3>
      </div>
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
        icon: 'pi pi-pencil',
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
            localStorage.clear();
            this.$router.push("/sign-in")
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

.title {
  margin-left: 150px;
  color: white;
}
</style>