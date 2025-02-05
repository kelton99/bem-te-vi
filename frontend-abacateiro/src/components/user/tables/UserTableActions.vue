<template>
  <q-card>
    <q-card-section>
      <q-table :rows="users" :columns="columns">
        <template v-slot:body-cell-actions="props">
          <q-td align="right" name="actions">
            <q-btn
              icon="edit"
              size="sm"
              class="q-ml-sm"
              flat
              dense
              @click="editUser(props.row)"
            />
            <q-btn
              icon="delete"
              size="sm"
              class="q-ml-sm"
              flat
              dense
              @click="deleteUser(props.row)"
            />
          </q-td>
        </template>
      </q-table>
    </q-card-section>

    <UserFormModal
      :is-modal-open="isModalOpen"
      :is-edit-mode="isEditMode"
      :user-data="selectedUser"
      @update:isModalOpen="isModalOpen = $event"
      @saveUser="handleSaveUser"
    />
  </q-card>
</template>

<script>
import { defineComponent, ref, onMounted, defineAsyncComponent } from "vue";
import { EventBus } from "@/plugins/eventBus";
import axios from "axios";

export default defineComponent({
  name: "TableActions",
  components: {
    UserFormModal: defineAsyncComponent(() =>
      import("components/user/form/UserFormModal.vue")
    ),
  },
  setup() {
    const users = ref([]);
    const columns = ref([
      {
        name: "user_name",
        required: true,
        label: "Name",
        align: "left",
        field: (row) => row.user_name,
        format: (val) => `${val}`,
        sortable: true,
      },
      {
        name: "user_email",
        required: true,
        label: "Email",
        align: "left",
        field: (row) => row.user_email,
        format: (val) => `${val}`,
        sortable: true,
      },
      {
        name: "user_document",
        label: "Document",
        align: "left",
        field: (row) => row.user_document,
        format: (val) => `${val}`,
        sortable: true,
      },
      { name: "actions", label: "Actions", align: "right" },
    ]);
    const isModalOpen = ref(false);
    const isEditMode = ref(true);

    const selectedUser = ref({
      user_id: 0,
      user_name: "",
      user_email: "",
      user_password: "",
      user_password_confirmation: "",
      user_document: "",
    });

    const fetchUsers = async () => {
      try {
        const response = await axios.get("http://localhost:8080/users");
        users.value = response.data;
      } catch (error) {
        console.error("Error fetching users:", error);
      }
    };

    const handleUserSaved = () => {
      fetchUsers();
    };

    onMounted(() => {
      fetchUsers();
      EventBus.on("user-saved", handleUserSaved);
    });

    const handleSaveUser = async (user) => {
      console.log("User to save:", user);

      if (isEditMode.value) {
        try {
          const response = await axios.put(
            `http://localhost:8080/users/${user.user_id}`,
            user
          );
          console.log("User updated successfully:", response.data);
          EventBus.emit("user-saved");
        } catch (error) {
          console.error("Error updating user:", error);
        }
      }
      closeModal();
    };

    const deleteUser = async (user) => {
      try {
        await axios.delete(`http://localhost:8080/users/${user.id}`);
        EventBus.emit("user-saved");
      } catch (error) {
        console.error("Error deleting user:", error);
      }
    };

    const openModal = () => {
      isModalOpen.value = true;
    };

    const closeModal = () => {
      isModalOpen.value = false;
    };

    const editUser = (user) => {
      selectedUser.value = {
        user_id: user.id,
        user_name: user.user_name,
        user_email: user.user_email,
        user_document: user.user_document,
      };
      openModal();
    };

    return {
      users,
      isModalOpen,
      isEditMode,
      selectedUser,
      columns,
      handleSaveUser,
      deleteUser,
      editUser,
    };
  },
});
</script>

<style scoped></style>
