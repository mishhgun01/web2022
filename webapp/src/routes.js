import MainView from "@/views/MainView";
import LoginComponent from "@/components/LoginComponent";

const routes = [
    {
        path: "/",
        component: MainView
    },
    {
        path: "/sign-in",
        component: LoginComponent
    }
]

export default routes