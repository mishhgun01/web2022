import MainView from "@/views/MainView";
import LoginComponent from "@/components/LoginComponent";
import NoteCard from "@/components/NoteCard";

const routes = [
    {
        path: "/",
        component: MainView
    },
    {
        path: "/sign-in",
        component: LoginComponent
    },
    {
        path: "/note-card/:opt",
        name: "noteCard",
        component: NoteCard
    }
]

export default routes