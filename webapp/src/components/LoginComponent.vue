<template>
    <div class="justify-content-center ml-100">
      <Toast  position="top-right"/>
        <Card style="width: 30rem; height: 30rem; margin-bottom: 5em">
            <template #header>
              {{ $t('auth') }}
            </template>
            <template #content>
                <div class="inner">
                    <span class="p-float-label inner">
                        <InputText v-model="login"  id="username" type="text"/>
                        <label v-if="!login.length" for="username">{{ $t('name') }}</label>
                    </span>
                    </div>
                <div class="inner">
                    <span class="p-float-label inner">
                        <Password v-model="password" toggleMask> </Password>
                        <label v-if="!password.length" for="pwd">{{ $t('pwd') }}</label>
                    </span>
                </div>
            </template>
            <template #footer>
              <ProgressSpinner v-if="spinner" style="width:60px;height:60px;margin-top: 65px" strokeWidth="6" fill="#FFFFFF" animationDuration=".7s"/>
                <div class="inner" v-if="!spinner">
                    <Button icon="pi pi-sign-in" iconPos="right" class="p-button-rounded p-button-raised" :label="$t('sign-in')" @click="onLogin"/>
                    <span class="inner">
                        <Button icon="pi pi-arrow-circle-right" iconPos="right" class="p-button-rounded p-button-link p-button-raised" :label="$t('reg')" @click="onRegister"/>
                    </span>
                </div>
            </template> 
        </Card>
    </div>
</template>
<script>
import Card from 'primevue/card/Card';
import Button from 'primevue/button/Button';
import InputText from 'primevue/inputtext';
import ProgressSpinner from 'primevue/progressspinner';
import Toast from 'primevue/toast';
import Password from 'primevue/password'
import { regUser, authUser, getUserInfo } from "@/api/user"
export default {
    components: {
        Card,
        Button,
        InputText,
        ProgressSpinner,
        Toast,
        Password
    },
    data() {
        return {
            spinner: false,
            login: "",
            password: ""
        }
    },
  methods: {
        onLogin() {
          this.spinner = true
          authUser(this.login, this.password)
            .then( () => {
              getUserInfo(this.login, this.password)
                .then(response=> {
                    const authToken = btoa(`${response.data.Login}:${response.data.Password}`)
                    localStorage.setItem("authToken", authToken)
                    localStorage.setItem("lastPath", response.data.LastPath)
                    localStorage.setItem("User", JSON.stringify(response.data))
                    this.$router.push(response.data.LastPath || "/")
                })
                .catch(() => {
                  this.$toast.add({
                    severity:'error',
                    summary: this.$t('error'),
                    detail: this.$t('server-error'),
                    life: 3000
                  });
                  this.spinner = false
              })
              .finally(()=>{
                this.spinner = false
              })
            })
            .catch(err => {
              this.spinner = false
                if (err.code === "ERR_BAD_REQUEST") {
                  this.$toast.add({
                    severity:'error',
                    summary: this.$t('error'),
                    detail: this.$t('login-error'),
                    life: 3000
                  });
                  return
                }

              this.$toast.add({
                severity:'error',
                summary: this.$t('error'),
                detail: this.$t('server-error'),
                life: 3000
              });
            })
        },
        onRegister() {
          this.spinner = true
            regUser(this.login, this.password)
            .then(r => {
                const authToken = btoa(`${r.data.Login}:${r.data.Password}`)
                localStorage.setItem("authToken", authToken)
                localStorage.setItem("User", JSON.stringify(r.data))
                this.spinner = false
                this.$router.push("/")
            })
            .catch(()=>{
                this.spinner = false
                this.$toast.add({
                  severity:'error',
                  summary: this.$t('error'),
                  detail: this.$t('server-error'),
                  life: 3000
                });
            })
            .finally(() => {
              this.spinner = false
            })
        }
    }
}
</script>
<style scoped>

.inner {
    display: flex;
    flex-direction: column;
    margin-top: 20px;
}

</style>