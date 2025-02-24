<template>
  <div class="container">
    <h1>Order Form</h1>
    <form @submit.prevent="submitForm">
      <div class="form-group">
        <label for="name">Name:</label>
        <input type="text" id="name" v-model="formData.name" placeholder="Enter Name" required />
      </div>

      <div class="form-group">
        <label for="email">Email:</label>
        <input type="email" id="email" v-model="formData.email" placeholder="Enter Email" required />
      </div>

      <div class="form-group">
        <label for="mobile">Mobile No.:</label>
        <input type="text" id="mobile" v-model="formData.mobileNo" placeholder="Enter Mobile No." required />
      </div>

      <div class="form-group">
        <label for="product">Product:</label>
        <select id="product" v-model="formData.product" required>
          <option value="">Select a product</option>
          <option value="Laptop">Laptop</option>
          <option value="Phone">Phone</option>
          <option value="Watch">Watch</option>
        </select>
      </div>

      <div class="form-group">
        <label for="quantity">Quantity:</label>
        <input type="number" id="quantity" v-model="formData.quantity" min="1" required />
      </div>

      <button type="submit">Submit Order</button>
    </form>
    <p v-if="message" class="p">{{ message }}</p>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'FORM',
  data() {
    return {
      formData: {
        name: '',
        email: '',
        mobileNo: '',
        product: '',
        quantity: ''
      },
      message: ''
    };
  },
  methods: {
    async submitForm() {
      try {
        const response = await axios.post('http://localhost:8000/api/register', this.formData);
        this.message = response.data.message;
        console.warn("User registered successfully!")
      } catch (error) {
        this.message = 'An error occurred while submitting the form.';
      }
    }
  }
};
</script>

<style>
.container {
  max-width: 400px;
  margin: auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 10px;
  box-shadow: 2px 2px 12px rgba(0, 0, 0, 0.1);
  background: #f9f9f9;
}

h1 {
  text-align: center;
  margin-bottom: 20px;
}

.form-group {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px;
}

label {
  flex: 1;
  font-weight: bold;
}

input,
select {
  flex: 2;
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

button {
  width: 100%;
  padding: 10px;
  background-color: #28a745;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 16px;
}

button:hover {
  background-color: #218838;
}

p {
  text-align: center;
  color: green;
  font-weight: bold;
}
</style>
