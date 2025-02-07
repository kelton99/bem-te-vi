import { route } from "quasar/wrappers";
import {
  createRouter,
  createMemoryHistory,
  createWebHistory,
  createWebHashHistory,
} from "vue-router";
import routes from "./routes";
import { useAuthStore } from "src/stores/authStore";

export default route(function (/* { store, ssrContext } */) {
  const createHistory = process.env.SERVER
    ? createMemoryHistory
    : process.env.VUE_ROUTER_MODE === "history"
    ? createWebHistory
    : createWebHashHistory;

  const Router = createRouter({
    scrollBehavior: () => ({ left: 0, top: 0 }),
    routes,
    history: createHistory(process.env.VUE_ROUTER_BASE),
  });

  // Adiciona uma guarda de rota global
  Router.beforeEach((to, from, next) => {
    const authStore = useAuthStore();

    // Carrega o token do localStorage se ainda não estiver carregado
    if (!authStore.token) {
      authStore.loadFromLocalStorage();
    }

    // Verifica se a rota requer autenticação
    if (to.matched.some((record) => record.meta.requiresAuth)) {
      // Se o usuário não estiver autenticado, redireciona para login
      if (!authStore.token) {
        next({ path: "/login" });
      } else {
        next(); // Se estiver autenticado, segue para a rota
      }
    } else {
      next(); // Se a rota não exigir autenticação, segue normalmente
    }
  });

  return Router;
});
