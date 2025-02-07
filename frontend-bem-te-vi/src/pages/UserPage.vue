<template>
  <q-page padding>
    <div class="button-container">
      <q-btn color="primary" label="Create User" @click="openCreateModal" />
    </div>

    <UserTableActions class="q-mt-lg"></UserTableActions>

    <UserFormModal
      :is-modal-open="isModalOpen"
      :is-edit-mode="isEditMode"
      :user-data="selectedUser"
      @update:isModalOpen="isModalOpen = $event"
      @saveUser="handleSaveUser"
    />
  </q-page>
</template>

<script>
import { ref, defineComponent, defineAsyncComponent } from "vue";
import axios from "axios";
import { EventBus } from "@/plugins/eventBus";

export default defineComponent({
  name: "UserPage",
  components: {
    UserTableActions: defineAsyncComponent(() =>
      import("components/user/tables/UserTableActions.vue")
    ),
    UserFormModal: defineAsyncComponent(() =>
      import("components/user/form/UserFormModal.vue")
    ),
  },
  setup() {
    const isModalOpen = ref(false);

    const isEditMode = ref(false);

    const selectedUser = ref({
      user_id: 0,
      user_name: "",
      user_email: "",
      user_password: "",
      user_password_confirmation: "",
      user_document: "",
    });

    const openModal = () => {
      isModalOpen.value = true;
    };

    const closeModal = () => {
      isModalOpen.value = false;
    };

    const openCreateModal = () => {
      selectedUser.value = {
        user_id: 0,
        user_name: "",
        user_email: "",
        user_password: "",
        user_password_confirmation: "",
        user_document: "",
      };
      openModal();
    };

    const handleSaveUser = async (user) => {
      if (!isEditMode.value) {
        try {
          const response = await axios.post(
            "http://localhost:8080/users",
            user
          );
          console.log("Usuário salvo com sucesso:", response.data);
          closeModal(); // Fechar modal
          EventBus.emit("user-saved"); // Emitir evento para atualizar UserTableActions
        } catch (error) {
          console.error("Error creating user:", error);
        }
      }
    };

    return {
      isModalOpen,
      isEditMode,
      selectedUser,
      openCreateModal,
      handleSaveUser,
    };
  },
});
</script>

<style scoped>
.button-container {
  display: flex;
  justify-content: flex-end;
}
</style>
