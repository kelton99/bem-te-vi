<template>
  <q-layout>
    <q-page-container>
      <q-page class="flex bg-image flex-center">
        <q-card
          v-bind:style="$q.screen.lt.sm ? { width: '80%' } : { width: '30%' }"
        >
          <q-card-section>
            <div class="text-center q-pt-lg">
              <div class="col text-h6 ellipsis">Log in</div>
            </div>
          </q-card-section>

          <q-card-section>
            <q-form @submit="handleSubmit" class="q-gutter-md">
              <q-input filled v-model="username" label="Username" lazy-rules />

              <q-input
                type="password"
                filled
                v-model="password"
                label="Password"
                lazy-rules
              />

              <div>
                <q-btn type="submit" label="Login" color="primary" />
              </div>
            </q-form>
          </q-card-section>
        </q-card>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script>
import { ref, defineComponent } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "src/stores/authStore";

export default defineComponent({
  setup() {
    const router = useRouter(); // Acessa o Vue Router
    const authStore = useAuthStore();

    const username = ref("");
    const password = ref("");

    const handleSubmit = async () => {
      const success = await authStore.login(username.value, password.value);

      if (success) {
        // Redirecionar para a página principal, por exemplo
        router.push("/");
      } else {
        // Exibir uma mensagem de erro ou tratar o erro adequadamente
        alert("Login falhou, tente novamente.");
      }
    };

    return {
      username,
      password,
      handleSubmit,
    };
  },
});
</script>

<style>
.bg-image {
  background-image: linear-gradient(135deg, #7028e4 0%, #e5b2ca 100%);
}
</style>
