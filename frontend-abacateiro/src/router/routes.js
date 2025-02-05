const routes = [
  {
    path: "/",
    component: () => import("layouts/MainLayout.vue"),
    children: [
      {
        path: "",
        component: () => import("pages/IndexPage.vue"),
        meta: { requiresAuth: true }, // Exige autenticação
      },
      {
        path: "/user",
        component: () => import("pages/UserPage.vue"),
        meta: { requiresAuth: true }, // Exige autenticação
      },
    ],
  },
  {
    path: "/login",
    component: () => import("pages/LoginAbcPage.vue"),
  },
  {
    path: "/:catchAll(.*)*",
    component: () => import("pages/ErrorNotFound.vue"),
  },
];

export default routes;
