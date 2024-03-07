<template>
  <div
      id="drop-area"
      @dragover.prevent="handleDragOver"
      @drop.prevent="handleFileDrop"
      @dragenter.prevent="dragEnter"
      @dragleave.prevent="dragLeave"
      class="drop-area"
  >
    <font-awesome-icon icon="cloud-upload-alt" size="5x"/>
    Drop files here to upload
    <!-- Modal -->
    <div v-if="showModal" class="modal" @click.self="showModal = false">
      <div class="modal-content">
        <span class="close" @click="showModal = false">&times;</span>
        <p>{{ uploadResponse }}</p>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import {library} from '@fortawesome/fontawesome-svg-core';
import {faCloudUploadAlt} from '@fortawesome/free-solid-svg-icons';
import {FontAwesomeIcon} from "@fortawesome/vue-fontawesome";

library.add(faCloudUploadAlt);

export default {
  components: {
    FontAwesomeIcon,
  },
  data() {
    return {
      showModal: false,
      uploadResponse: '',
    };
  },
  methods: {
    handleDragOver(event) {
      event.preventDefault();
    },
    dragEnter(event) {
      event.target.style.border = "2px solid #000"; // Visual feedback
    },
    dragLeave(event) {
      event.target.style.border = ""; // Revert visual feedback
    },
    handleFileDrop(event) {
      const files = event.dataTransfer.files;
      this.uploadFiles(files[0]); // Assuming single file upload, adjust as needed
      event.target.style.border = ""; // Revert visual feedback
    },
    async uploadFiles(file) {
      const formData = new FormData();
      formData.append('file', file);

      try {
        const response = await axios.post('http://localhost:8082/api/upload', formData, {
          headers: {
            'Content-Type': 'multipart/form-data',
          },
        });
        console.log(response.data); // Handle success
        if (response.status === 200) {
          // pretty modal
          this.uploadResponse = response.data;
          this.showModal = true; // Show the modal with the result
        }
      } catch (error) {
        console.error(error); // Handle error
        this.uploadResponse = error.response.data;
        this.showModal = true; // Optionally, show the modal with error
      }
    },
  },
};
</script>

<style scoped>
.drop-area {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 20px; /* Adjust the space between the icon and the text */
  border: 2px dashed #ccc;
  border-radius: 5px;
  width: 100%; /* Adjust based on your preference or use max-width for responsiveness */
  min-width: 300px; /* Ensures it doesn't get too small on narrow screens */
  height: 200px; /* Adjust based on your preference */
  text-align: center;
  padding: 20px;
  box-sizing: border-box;
  transition: border-color 0.3s;
}

.drop-area:hover {
  border-color: #000; /* Optional: Change border color on hover */
}

.modal {
  position: fixed;
  z-index: 1;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  overflow: auto;
  background-color: rgb(0, 0, 0);
  background-color: rgba(0, 0, 0, 0.4);
}

.modal-content {
  background-color: #fefefe;
  margin: 15% auto;
  padding: 20px;
  border: 1px solid #888;
  width: 80%;
}

.close {
  color: #aaa;
  float: right;
  font-size: 28px;
  font-weight: bold;
}

.close:hover,
.close:focus {
  color: black;
  text-decoration: none;
  cursor: pointer;
}
</style>
