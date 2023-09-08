<template>
    <div class="header">
        <div class="header-btns">
            <SpeedDial :model="items"/>
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
            this.$toast.add({ severity: 'info', summary: 'Add', detail: 'Data Added' });
        }
    },
    {
        label: 'Update',
        icon: 'pi pi-refresh',
        command: () => {
            this.$toast.add({ severity: 'success', summary: 'Update', detail: 'Data Updated' });
        }
    },
    {
        label: 'Delete',
        icon: 'pi pi-trash',
        command: () => {
            this.$toast.add({ severity: 'error', summary: 'Delete', detail: 'Data Deleted' });
        }
    },
    {
        label: 'Upload',
        icon: 'pi pi-upload',
        command: () => {
            window.location.hash = "/fileupload"
        }
    },
    {
        label: 'Vue Website',
        icon: 'pi pi-external-link',
        command: () => {
            window.location.href = 'https://vuejs.org/'
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
    display: flex;
    flex-direction: row;
    align-content: space-between;
    background-color: #94d1be;
    margin: 0;
    padding: 5px;
    width: 100%;
    height: 100px;
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