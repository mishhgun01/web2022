import Vue from 'vue'
import VueI18n from "vue-i18n";

Vue.use(VueI18n)

const words = {
    ru: {
        'auth': 'Авторизация',
        'sign-in': 'Войти',
        'reg': 'Зарегистрироваться',
        'name': 'Имя',
        'pwd': 'Пароль',
        'delete-note': 'Удалить заметку',
        'sure-text': 'Вы уверены?',
        'yes': 'Да',
        'no': 'Нет',
        'sign-out': 'Выйти',
        'save': 'Сохранить',
        'create': 'Создать',
        'delete':'Удалить',
        'edit': 'Редактировать',
        'my-notes': 'Мои заметки',
        'no-notes-text': 'Ни одной заметки!',
        'create-sugg': 'Создайте новую заметку!',
        'note-title': 'Название',
        'status_text': 'Статус',
        'status_1': 'Начато 📌',
        'status_2': 'В работе 💻',
        'status_3': 'Выполнено ✅',
        'desc': 'Описание',
        'error': 'Ошибка',
        'note-error': 'Необходимо заполнить название и описание.',
        'login-error': 'Неверное имя пользователя или пароль.',
        'server-error': 'Сервер не отвечает. Попробуйте позже.'
    },
    en: {
        'auth': 'Authorization',
        'sign-in': 'Sign in',
        'reg': 'Registration',
        'name': 'Name',
        'pwd': 'Password',
        'delete-note': 'Delete note',
        'delete': 'Delete',
        'sure-text': 'Are you sure?',
        'yes': 'Yes',
        'no': 'No',
        'sign-out': 'Sign out',
        'save': 'Save',
        'create': 'Create',
        'edit': 'Edit',
        'no-notes-text': 'No one note!',
        'create-sugg': 'Create a new one!',
        'note-title': 'Title',
        'my-notes': 'My notes',
        'status_text': 'Status',
        'status_1': 'Started 📌',
        'status_2': 'In progress 💻',
        'status_3': 'Done ✅',
        'desc': 'Description',
        'error': 'Error',
        'note-error': 'There must be something in Title and Description.',
        'login-error': 'Wrong name or password.',
        'server-error': 'Server is temporary unavailable.Try again later.'
    }
}

const locale = localStorage.getItem("locale") || "en"

export const i18n = new VueI18n({
    locale,
    fallbackLocale: 'en',
    messages: words
})