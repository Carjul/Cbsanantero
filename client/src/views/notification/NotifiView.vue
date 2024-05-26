<template>
    <div class="container">
        <br>
 <!-- no toques este codigo  -->

<div data-bs-toggle="modal" data-bs-target="#messageModal" class="shadow-lg p-3 mb-5 bg-body-tertiary rounded message">
    <i class="fas fa-envelope message-icon"></i>
    <div class="message-content">Alfredo Gonzales</div>
  </div>

  <!-- Modal -->
  <div class="modal fade" id="messageModal" tabindex="-1" aria-labelledby="messageModalLabel" aria-hidden="true">
    <div class="modal-dialog modal-lg">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="messageModalLabel">Solicitud</h5>
          <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
        </div>
        <div class="modal-body">
          <p>Aquí puedes colocar el contenido completo del mensaje. Este es un ejemplo de cómo se vería un mensaje más largo y detallado.</p>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Cerrar</button>
        </div>
      </div>
    </div>
  </div>


    </div>

   
</template>

<style>
 .message {
      position: relative;
      display: flex;
      align-items: center;
      padding: 20px;
      margin-bottom: 20px;
      border-radius: 10px;
      box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
      background-color: #f8f9fa; /* bg-body-tertiary */
      transition: background-color 0.3s ease;
    }

    .message:hover {
      background-color: #e2e6ea;
    }

    .message-icon {
      position: relative;
      font-size: 24px;
      margin-right: 15px;
      color: #dc3545; /* bg-danger */
    }

    .message-icon::after {
      content: '';
      position: absolute;
      top: -5px;
      right: -5px;
      width: 10px;
      height: 10px;
      background-color: #28a745; /* Green indicator */
      border-radius: 50%;
      border: 2px solid white; /* Optional: border to make the dot stand out */
    }

    .message-content {
      font-size: 18px;
    }
</style>