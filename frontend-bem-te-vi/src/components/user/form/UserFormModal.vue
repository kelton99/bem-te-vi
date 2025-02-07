<template>
  <q-dialog v-model="localIsModalOpen" @hide="onDialogHide">
    <q-card
      v-bind:style="$q.screen.lt.sm ? { width: '80%' } : { width: '40%' }"
    >
      <q-card-section>
        <div class="text-center q-pt-lg">
          <div class="col text-h6 ellipsis">
            {{ isEditMode ? "Edit User" : "Create User" }}
          </div>
        </div>
      </q-card-section>
      <q-card-section>
        <q-form class="q-gutter-md">
          <div class="col-6">
            <q-item>
              <q-input
                dense
                outlined
                class="full-width"
                v-model="user.user_name"
                label="Nome"
                :error="v$.user_name.$invalid && v$.user_name.$dirty"
                :error-message="v$.user_name.$invalid ? 'Nome inválido' : ''"
              />
            </q-item>
          </div>
          <div class="col-6">
            <q-item>
              <q-input
                dense
                outlined
                class="full-width"
                v-model="user.user_email"
                label="Email"
                :error="v$.user_email.$invalid && v$.user_email.$dirty"
                :error-message="v$.user_email.$invalid ? 'Email inválido' : ''"
              />
            </q-item>
          </div>
          <div class="col-6">
            <q-item>
              <q-input
                dense
                outlined
                class="full-width"
                v-model="user.user_document"
                label="Documento"
                :error="v$.user_document.$invalid && v$.user_document.$dirty"
                :error-message="
                  v$.user_document.$invalid ? 'Documento inválido' : ''
                "
              />
            </q-item>
          </div>
          <div class="col-6" v-if="!isEditMode">
            <q-item>
              <q-input
                dense
                outlined
                class="full-width"
                v-model="user.user_password"
                label="Senha"
                type="password"
                :error="v$.user_password.$invalid && v$.user_password.$dirty"
                :error-message="
                  v$.user_password.$invalid ? 'Senha inválida' : ''
                "
              />
            </q-item>
          </div>
          <div class="col-6" v-if="!isEditMode">
            <q-item>
              <q-input
                dense
                outlined
                class="full-width"
                v-model="user.user_password_confirmation"
                label="Confirme a Senha"
                type="password"
                :error="
                  v$.user_password_confirmation.$invalid &&
                  v$.user_password_confirmation.$dirty
                "
                :error-message="
                  v$.user_password_confirmation.$invalid
                    ? 'As senhas não coincidem'
                    : ''
                "
              />
            </q-item>
          </div>
          <div class="col-6">
            <q-card-actions align="right">
              <q-btn label="Fechar" color="primary" @click="closeModal" />
              <q-btn label="Salvar" color="primary" @click="saveUser" />
            </q-card-actions>
          </div>
        </q-form>
      </q-card-section>
    </q-card>
  </q-dialog>
</template>

<script>
import { ref, computed, watch, defineComponent } from "vue";
import useVuelidate from "@vuelidate/core";
import { required, email, minLength, sameAs } from "@vuelidate/validators";

export default defineComponent({
  name: "UserModal",
  props: {
    isEditMode: {
      type: Boolean,
      default: false,
    },
    isModalOpen: {
      type: Boolean,
      required: true,
    },
    userData: {
      type: Object,
      default: () => ({
        user_id: 0,
        user_name: "",
        user_email: "",
        user_password: "",
        user_password_confirmation: "",
        user_document: "",
      }),
    },
  },
  emits: ["update:isModalOpen", "saveUser"],
  setup(props, { emit }) {
    const user = ref({ ...props.userData });
    const localIsModalOpen = ref(props.isModalOpen);
    const password = computed(() => user.value.user_password);

    watch(
      () => props.isModalOpen,
      (newVal) => {
        localIsModalOpen.value = newVal;
      }
    );

    watch(
      () => props.userData,
      (newVal) => {
        user.value = { ...newVal };
      }
    );

    const rules = computed(() => {
      const baseRules = {
        user_name: { required },
        user_email: { required, email },
        user_document: { required },
      };

      if (props.isEditMode) {
        return baseRules;
      }

      return {
        ...baseRules,
        user_password: { required, minLength: minLength(6) },
        user_password_confirmation: {
          required,
          sameAsPassword: sameAs(password),
        },
      };
    });

    /**
     * @TODO
     * Sempre que o saveUser é chamado, preciso de uma instância nova do v$.
     */
    const v$ = useVuelidate(rules, user);

    const closeModal = () => {
      localIsModalOpen.value = false;

      user.value = {
        user_id: 0,
        user_name: "",
        user_email: "",
        user_password: "",
        user_password_confirmation: "",
        user_document: "",
      };

      v$.value.$reset();

      emit("update:isModalOpen", false);
    };

    const onDialogHide = () => {
      closeModal();
    };

    const saveUser = () => {

      /**
       * @TODO
       * Veficar poque na segunda vez que eu clico no botão salvar, da erro.
       * Acredito que seja por causa do $touch tentade fazer o $reset() antes do $touch()
       */
      v$?.value?.$touch();

      if (v$?.value?.$invalid) {
        console.log("Invalid form");
        return;
      }

      emit("saveUser", user.value);

      closeModal();
    };

    return {
      user,
      localIsModalOpen,
      closeModal,
      saveUser,
      onDialogHide,
      v$,
    };
  },
});
</script>
