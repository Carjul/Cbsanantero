<template>
    <div class="home">
      <!--   <video autoplay muted loop id="video-background">
            <source src="https://videocdn.cdnpk.net/joy/content/video/free/video0467/large_preview/_import_61501399ee3f05.37978127.mp4?filename=1109422_1080p_4k_2k_3840x2160.mp4" type="video/mp4">
            Tu navegador no soporta video HTML5.
            </video> -->



        <!-- Carousel Start -->
        <div class="container-fluid p-0">

            <div id="header-carousel" class="carousel slide" data-ride="carousel">
                <div class="carousel-inner">
                    <div class="carousel-item active">
                        <img class="w-100"
                            src="https://i.pinimg.com/736x/2e/e0/65/2ee06566abc4d0cfc5ef608919fbe1e8.jpg"
                            alt="Image">
                        <div class="carousel-caption d-flex flex-column align-items-center justify-content-center">
                            <div class="p-3" style="max-width: 900px;">
                                <h4 class="text-white text-uppercase mb-md-3">¡En hora buena!</h4>
                                <h1 class="display-3 text-white mb-md-4">Bienvenido a Costa Brisa</h1>

                                <router-link to="/hotels" class="btn btn-primary py-md-3 px-md-5 mt-2 btn-book-now">
                                    Solicitar Servicio
                                    </router-link>


                            </div>
                        </div>
                    </div>
                    <div class="carousel-item">
                        <img class="w-100"
                            src="https://i.pinimg.com/736x/63/d7/a3/63d7a3fbc7105ad21bbd015305f9d98f.jpg"
                            alt="Image">
                        <div class="carousel-caption d-flex flex-column align-items-center justify-content-center">
                            <div class="p-3" style="max-width: 900px;">
                                <h4 class="text-white text-uppercase mb-md-3">HORA DE CONOCER</h4>
                                <h1 class="display-3 text-white mb-md-4">Descubre Lugares Increíbles Con Nosotros</h1>
                                <router-link to="/hotels" class="btn btn-primary py-md-3 px-md-5 mt-2 btn-book-now">
                                    Solicitar Servicio
                                    </router-link>

                            </div>
                        </div>
                    </div>
                </div>
                <a class="carousel-control-prev" href="#header-carousel" data-slide="prev">
                    <div class="btn btn-dark" style="width: 45px; height: 45px;">
                        <span class="carousel-control-prev-icon mb-n2"></span>
                    </div>
                </a>
                <a class="carousel-control-next" href="#header-carousel" data-slide="next">
                    <div class="btn btn-dark" style="width: 45px; height: 45px;">
                        <span class="carousel-control-next-icon mb-n2"></span>
                    </div>
                </a>
            </div>
        </div>
        <!-- Carousel End -->







        <!-- Booking Start --
    <div class="container-fluid booking mt-5 pb-5">
        <div class="container pb-5">
            <div class="bg-light shadow" style="padding: 30px;">
                <div class="row align-items-center" style="min-height: 60px;">
                    <div class="col-md-10">
                        <div class="row">
                            <div class="col-md-3">
                                <div class="mb-3 mb-md-0">
                                    <select class="custom-select px-4" style="height: 47px;">
                                        <option selected>Destination</option>
                                        <option value="1">Destination 1</option>
                                        <option value="2">Destination 1</option>
                                        <option value="3">Destination 1</option>
                                    </select>
                                </div>
                            </div>
                            <div class="col-md-3">
                                <div class="mb-3 mb-md-0">
                                    <div class="date" id="date1" data-target-input="nearest">
                                        <input type="text" class="form-control p-4 datetimepicker-input" placeholder="Depart Date" data-target="#date1" data-toggle="datetimepicker"/>
                                    </div>
                                </div>
                            </div>
                            <div class="col-md-3">
                                <div class="mb-3 mb-md-0">
                                    <div class="date" id="date2" data-target-input="nearest">
                                        <input type="text" class="form-control p-4 datetimepicker-input" placeholder="Return Date" data-target="#date2" data-toggle="datetimepicker"/>
                                    </div>
                                </div>
                            </div>
                            <div class="col-md-3">
                                <div class="mb-3 mb-md-0">
                                    <select class="custom-select px-4" style="height: 47px;">
                                        <option selected>Duration</option>
                                        <option value="1">Duration 1</option>
                                        <option value="2">Duration 1</option>
                                        <option value="3">Duration 1</option>
                                    </select>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="col-md-2">
                        <button class="btn btn-primary btn-block" type="submit" style="height: 47px; margin-top: -2px;">Submit</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
     Booking End -->
        <hr>

        <!-- About Start -->
        <div class="container-fluid py-6">
            <div class="container pt-6">
                <div class="row">
                    <div class="col-lg-6 position-relative" style="min-height: 500px;">
                        <div
                            class="position-absolute h-100 top-0 bottom-0 left-0 right-0 d-flex align-items-center justify-content-center">
                            <img class="w-100 h-100"
                                src="https://i.pinimg.com/736x/94/29/d0/9429d017040c47a6e29c4559e4526ba4.jpg "
                                style="object-fit: cover;">
                        </div>
                    </div>
                    <div class="col-lg-6 pt-5 pb-lg-5">
                        <div class="about-text bg-white p-4 p-lg-5 my-lg-5">
                            <h6 class="text-primary text-uppercase" style="letter-spacing: 5px;"> SOBRE NOSOTROS</h6>
                            <h1 class="mb-3">Ofrecemos los mejores alojamientos en la Costa Caribe Colombiana.</h1>
                            <p>

                                Desde la selección de destinos hasta la elección de alojamientos y actividades, estamos
                                comprometidos a brindarte una experiencia que supere tus expectativas.

                            </p>
                            <div class="row mb-4">
                                <div class="col-6">
                                    <img class="img-fluid"
                                        src="https://i.pinimg.com/736x/58/b7/82/58b782c63413df5977594a92a275ac04.jpg"
                                        alt="">
                                </div>
                                <div class="col-6">
                                    <img class="img-fluid"
                                        src="https://i.pinimg.com/736x/17/bf/4a/17bf4a295745963b32a5009ee018e31d.jpg"
                                        alt="">
                                </div>
                            </div>
                            <router-link to="/hotels" class="btn btn-primary py-md-3 px-md-5 mt-2 btn-book-now">
                                    Solicitar Servicio
                                    </router-link>

                        </div>
                    </div>
                </div>
            </div>
        </div>
        <!-- About End -->




        <hr>







        <!--  <div class="container-fluid pb-5">
        <div class="container pb-5">
            <div class="row">
                <div class="col-md-4">
                    <div class="d-flex mb-4 mb-lg-0">
                        <div class="d-flex flex-shrink-0 align-items-center justify-content-center bg-primary mr-3" style="height: 100px; width: 100px;">
                            <i class="fa fa-2x fa-money-check-alt text-white"></i>
                        </div>
                        <div class="d-flex flex-column">
                            <h5 class="">Competitive Pricing</h5>
                            <p class="m-0">Magna sit magna dolor duo dolor labore rebum amet elitr est diam sea</p>
                        </div>
                    </div>
                </div>
                <div class="col-md-4">
                    <div class="d-flex mb-4 mb-lg-0">
                        <div class="d-flex flex-shrink-0 align-items-center justify-content-center bg-primary mr-3" style="height: 100px; width: 100px;">
                            <i class="fa fa-2x fa-award text-white"></i>
                        </div>
                        <div class="d-flex flex-column">
                            <h5 class="">Best Services</h5>
                            <p class="m-0">Magna sit magna dolor duo dolor labore rebum amet elitr est diam sea</p>
                        </div>
                    </div>
                </div>
                <div class="col-md-4">
                    <div class="d-flex mb-4 mb-lg-0">
                        <div class="d-flex flex-shrink-0 align-items-center justify-content-center bg-primary mr-3" style="height: 100px; width: 100px;">
                            <i class="fa fa-2x fa-globe text-white"></i>
                        </div>
                        <div class="d-flex flex-column">
                            <h5 class="">Worldwide Coverage</h5>
                            <p class="m-0">Magna sit magna dolor duo dolor labore rebum amet elitr est diam sea</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
     -->


        <!--   Destinos  -->

        <div class="container-fluid py-5">
            <div class="container pt-5 pb-3">
                <div class="text-center mb-3 pb-3">
                    <h6 class="text-primary text-uppercase" style="letter-spacing: 5px;">Hoteles</h6>
                    <h1>Destacados</h1>
                </div>
                <div class="row">
                    <div class="col-lg-4 col-md-6 mb-4">
                        <div class="destination-item position-relative overflow-hidden mb-2">
                            <img class="img-fluid"
                            style="width:500px;height:250px;"        src="https://i.pinimg.com/736x/63/d7/a3/63d7a3fbc7105ad21bbd015305f9d98f.jpg"
                                alt="">
                                
                                <router-link class="destination-overlay text-white text-decoration-none" to="/hotels">
                                    <h5 class="text-white">Hotel Dubai</h5>
                                    <span>San Antero</span>
                                    </router-link>

                        </div>
                    </div>
                    <div class="col-lg-4 col-md-6 mb-4">
                        <div class="destination-item position-relative overflow-hidden mb-2">
                            <img class="img-fluid"
                            style="width:500px;height:250px;"      src="https://i.pinimg.com/736x/2e/e0/65/2ee06566abc4d0cfc5ef608919fbe1e8.jpg"
                                alt="">
                                <router-link class="destination-overlay text-white text-decoration-none" to="/hotels">
                                <h5 class="text-white">Hotel Maria Mulata</h5>
                                <span>San Antero</span>
                                </router-link>
                        </div>
                    </div>
                    <div class="col-lg-4 col-md-6 mb-4">
                        <div class="destination-item position-relative overflow-hidden mb-2">
                            <img class="img-fluid"
                            style="width:500px;height:250px;"    src="https://i.pinimg.com/736x/74/2a/6b/742a6be53be5c8122476976f51fb56ef.jpg"
                                alt="">
                                <router-link class="destination-overlay text-white text-decoration-none" to="/hotels">
                                <h5 class="text-white">Hotel Perlayana</h5>
                                <span>San Antero</span>
                                </router-link>
                        </div>
                    </div>
                    <!-- <div class="col-lg-4 col-md-6 mb-4">
                        <div class="destination-item position-relative overflow-hidden mb-2">
                            <img class="img-fluid"
                                src="https://img.freepik.com/foto-gratis/piscina-playa-hotel-lujo-tipo-complejo-entretenimiento-amara-dolce-vita-hotel-lujo-recurso-tekirova-kemer-pavo_146671-18726.jpg?t=st=1710034827~exp=1710038427~hmac=f0c1f5ba02a73c299f6ec4abf2f74f267f5a9ea8940ec36d1b7808c818030600&w=900"
                                alt="">
                            <a class="destination-overlay text-white text-decoration-none" href="">
                                <h5 class="text-white">El Tonner</h5>
                                <span>120 Visitas</span>
                            </a>
                        </div>
                    </div> -->
                    <!-- <div class="col-lg-4 col-md-6 mb-4">
                        <div class="destination-item position-relative overflow-hidden mb-2">
                            <img class="img-fluid"
                                src="https://img.freepik.com/foto-gratis/tipo-complejo-entretenimiento-popular-complejo-piscinas-parques-acuaticos-turquia-hotel-lujo-recurso_146671-18827.jpg?t=st=1710035820~exp=1710039420~hmac=758c1a330afa4df6e036a65708c4252b9b6e2d5b1670c726752740c59ca49bad&w=900"
                                alt="">
                            <a class="destination-overlay text-white text-decoration-none" href="">
                                <h5 class="text-white">Condominio</h5>
                                <span>300 Visitas</span>
                            </a>
                        </div>
                    </div> -->
                    <!-- <div class="col-lg-4 col-md-6 mb-4">
                        <div class="destination-item position-relative overflow-hidden mb-2">
                            <img class="img-fluid"
                                src="https://img.freepik.com/foto-gratis/hermosa-piscina-lujo-al-aire-libre-hotel-resort_74190-7433.jpg?t=st=1710034886~exp=1710038486~hmac=ec17208b54994f9a9ca75cf0f1315a49ea314572bcc4c57ef737aa1c103cc5f9&w=900"
                                alt="">
                            <a class="destination-overlay text-white text-decoration-none" href="">
                                <h5 class="text-white">La Playa</h5>
                                <span>200 Visitas</span>
                            </a>
                        </div>
                    </div> -->
                </div>
            </div>
        </div>


        <!-- fin destinos  -->


        <!-- Service Start -->
        <!--  <div class="container-fluid py-6">
        <div class="container pt-6 pb-3">
            <div class="text-center mb-3 pb-3">
                <h6 class="text-primary text-uppercase" style="letter-spacing: 5px;">Services</h6>
                <h1>hh</h1>
            </div>
            <div class="row">
                <div class="col-lg-4 col-md-6 mb-4">
                    <div class="service-item bg-white text-center mb-2 py-5 px-4">
                        <i class="fa fa-2x fa-route mx-auto mb-4"></i>
                        <h5 class="mb-2">Travel Guide</h5>
                        <p class="m-0">Justo sit justo eos amet tempor amet clita amet ipsum eos elitr. Amet lorem est amet labore</p>
                    </div>
                </div>
                <div class="col-lg-4 col-md-6 mb-4">
                    <div class="service-item bg-white text-center mb-2 py-5 px-4">
                        <i class="fa fa-2x fa-ticket-alt mx-auto mb-4"></i>
                        <h5 class="mb-2">Ticket Booking</h5>
                        <p class="m-0">Justo sit justo eos amet tempor amet clita amet ipsum eos elitr. Amet lorem est amet labore</p>
                    </div>
                </div>
                <div class="col-lg-4 col-md-6 mb-4">
                    <div class="service-item bg-white text-center mb-2 py-5 px-4">
                        <i class="fa fa-2x fa-hotel mx-auto mb-4"></i>
                        <h5 class="mb-2">Hotel Booking</h5>
                        <p class="m-0">Justo sit justo eos amet tempor amet clita amet ipsum eos elitr. Amet lorem est amet labore</p>
                    </div>
                </div>
            </div>
        </div>
    </div> -->
        <!-- Service End -->


        <!-- Packages Start -->

        <div class="container-fluid py-6">
            <div class="container pt-6 pb-3">

                <div class="text-center mb-3 pb-3">
                    <h6 class="text-primary text-uppercase" style="letter-spacing: 5px;">Nuestros</h6>
                    <h1>Hoteles</h1>
                </div>
                <div class="row">
                    <div class="col-lg-4 col-md-6 mb-4">
  <div class="package-item bg-white mb-2">
    <img class="img-fluid" src="https://i.pinimg.com/736x/74/2a/6b/742a6be53be5c8122476976f51fb56ef.jpg" alt="">
    <div class="p-4">
      <div class="d-flex justify-content-between mb-3">
        <small class="m-0"><i class="fa fa-map-marker-alt text-primary mr-2"></i>San Antero</small>
        <small class="m-0"><i class="fa fa-hotel text-primary mr-2"></i>Perlayana</small>
      </div>
      <router-link class="h5 text-decoration-none" to="/hotels">
        El condominio cuenta con piscina, vigilancia privada..
      </router-link>
      <div class="border-top mt-4 pt-4">
        <div class="d-flex justify-content-between">
          <h6 class="m-0"><i class="fa fa-star text-primary mr-2"></i>4.5
            <small>(200)</small>
          </h6>
          <h5 class="m-0">$60.000</h5>
        </div>
      </div>
    </div>
  </div>
</div>

                    <div class="col-lg-4 col-md-6 mb-4">
                        <div class="package-item bg-white mb-2">
                            <img class="img-fluid"
                                src="https://i.pinimg.com/736x/63/d7/a3/63d7a3fbc7105ad21bbd015305f9d98f.jpg"
                                alt="">
                            <div class="p-4">
                                <div class="d-flex justify-content-between mb-3">
                                    <small class="m-0"><i
                                            class="fa fa-map-marker-alt text-primary mr-2"></i>San Antero</small>
                                   
                                    <small class="m-0"><i class="fa fa-hotel text-primary mr-2"></i>Dubai</small>
                                </div>
                                <router-link class="h5 text-decoration-none" to="/hotels">
                                    Ofrece piscina al aire libre, jardín, terraza y restaurante..
                                </router-link>
                                
                                <div class="border-top mt-4 pt-4">
                                    <div class="d-flex justify-content-between">
                                        <h6 class="m-0"><i class="fa fa-star text-primary mr-2"></i>4.5
                                            <small>(150)</small>
                                        </h6>
                                        <h5 class="m-0">$80.000</h5>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="col-lg-4 col-md-6 mb-4">
                        <div class="package-item bg-white mb-2">
                            <img class="img-fluid"
                                src="https://i.pinimg.com/736x/63/82/be/6382befaf24cf7e1902390695025e72d.jpg"
                                alt="">
                            <div class="p-4">
                                <div class="d-flex justify-content-between mb-3">
                                    <small class="m-0"><i
                                            class="fa fa-map-marker-alt text-primary mr-2"></i>San Antero</small>
                                 
                                    <small class="m-0"><i class="fa fa-hotel text-primary mr-2"></i>Camping Cispata</small>
                                </div>
                                <router-link class="h5 text-decoration-none" to="/hotels">Ofrece un hermoso mar, playa, sol y un bello paisaje de árboles..</router-link>
                                
                                <div class="border-top mt-4 pt-4">
                                    <div class="d-flex justify-content-between">
                                        <h6 class="m-0"><i class="fa fa-star text-primary mr-2"></i>4.5
                                            <small>(180)</small>
                                        </h6>
                                        <h5 class="m-0">$70.000</h5>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="col-lg-4 col-md-6 mb-4">
                        <div class="package-item bg-white mb-2">
                            <img class="img-fluid"
                                src="https://i.pinimg.com/736x/f9/08/39/f90839c2030ab2ff4b06462c7485c82d.jpg"
                                alt="">
                            <div class="p-4">
                                <div class="d-flex justify-content-between mb-3">
                                    <small class="m-0"><i
                                            class="fa fa-map-marker-alt text-primary mr-2"></i>San Antero</small>
                                   
                                    <small class="m-0"><i class="fa fa-hotel text-primary mr-2"></i>Apartahotel</small>
                                </div>
                                <router-link class="h5 text-decoration-none" to="/hotels">El hotel cuenta con piscina,servicio de Wi-Fi gratuito..</router-link>
                                <div class="border-top mt-4 pt-4">
                                    <div class="d-flex justify-content-between">
                                        <h6 class="m-0"><i class="fa fa-star text-primary mr-2"></i>4.5
                                            <small>(160)</small>
                                        </h6>
                                        <h5 class="m-0">$60.000</h5>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="col-lg-4 col-md-6 mb-4">
                        <div class="package-item bg-white mb-2">
                            <img class="img-fluid"
                                src="https://i.pinimg.com/736x/37/f6/84/37f6841ad68cb94e29489076fae8e59c.jpg"
                                alt="">
                            <div class="p-4">
                                <div class="d-flex justify-content-between mb-3">
                                    <small class="m-0"><i
                                            class="fa fa-map-marker-alt text-primary mr-2"></i>San Antero</small>
                                   
                                    <small class="m-0"><i class="fa fa-hotel text-primary mr-2"></i>Mono Cuco</small>
                                </div>
                                <router-link class="h5 text-decoration-none" to="/hotels">Ofrece alojamiento con salón compartido, estacionamiento..</router-link>
                                <div class="border-top mt-4 pt-4">
                                    <div class="d-flex justify-content-between">
                                        <h6 class="m-0"><i class="fa fa-star text-primary mr-2"></i>4.5
                                            <small>(190)</small>
                                        </h6>
                                        <h5 class="m-0">$90.000</h5>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="col-lg-4 col-md-6 mb-4">
                        <div class="package-item bg-white mb-2">
                            <img class="img-fluid"
                                src="https://i.pinimg.com/736x/58/a0/e8/58a0e8250890b1cc253525f720c12587.jpg"
                                alt="">
                            <div class="p-4">
                                <div class="d-flex justify-content-between mb-3">
                                    <small class="m-0"><i
                                            class="fa fa-map-marker-alt text-primary mr-2"></i>San Antero</small>
                                 
                                    <small class="m-0"><i class="fa fa-user text-primary mr-2"></i>Uvita de Playa</small>
                                </div>
                                <router-link class="h5 text-decoration-none" to="/hotels">El hotel ofrece cómodas habitaciones con aire..</router-link>
                                <div class="border-top mt-4 pt-4">
                                    <div class="d-flex justify-content-between">
                                        <h6 class="m-0"><i class="fa fa-star text-primary mr-2"></i>4.5
                                            <small>(250)</small>
                                        </h6>
                                        <h5 class="m-0">$40.000</h5>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <!-- Packages End -->

        <div class="container-fluid py-6">
            <div class="container pt-6 pb-3">
                <div class="text-center mb-3 pb-3">
                    <h6 class="text-primary text-uppercase" style="letter-spacing: 5px;">Nuestros</h6>
                    <h1>Restaurantes</h1>
                </div>
                <div class="row pb-3">
                    <div class="col-lg-4 col-md-6 mb-4 pb-2">
                        <div class="blog-item">
                            <div class="position-relative">
                                <img class="img-fluid w-100"
                                style="width:500px;height:250px;"    src="https://i.pinimg.com/736x/c7/1d/75/c71d754462be0ebea7cf040d8836832f.jpg"
                                    alt="">
                                <div class="blog-date">
                                    <h6 class="font-weight-bold mb-n1">Costa</h6>
                                    <small class="text-white text-uppercase">Brisa</small>
                                </div>
                            </div>
                            <div class="bg-white p-4">
                                <div class="d-flex mb-2">
                                    <!--  <a class="text-primary text-uppercase text-decoration-none" href="">Admin</a>
                                <span class="text-primary px-2">|</span> -->
                                    <!--  <a class="text-primary text-uppercase text-decoration-none" href="">Tours & Travel</a> -->
                                </div>
                                
                                <router-link class="h5 m-0 text-decoration-none" to="/restaurantes">
                                Restaurante Cangrejo Azul
                                </router-link>

                            </div>
                        </div>
                    </div>
                    <div class="col-lg-4 col-md-6 mb-4 pb-2">
                        <div class="blog-item">
                            <div class="position-relative">
                                <img class="img-fluid w-100"
                                style="width:500px;height:250px;"  src="https://i.pinimg.com/736x/f4/59/ee/f459ee1e439f4473743a963598bcdc72.jpg"
                                    alt="">
                                <div class="blog-date">
                                    <h6 class="font-weight-bold mb-n1">Costa</h6>
                                    <small class="text-white text-uppercase">Brisa</small>
                                </div>
                            </div>
                            <div class="bg-white p-4">
                                <div class="d-flex mb-2">
                                    <!--  <a class="text-primary text-uppercase text-decoration-none" href="">Admin</a>
                                <span class="text-primary px-2">|</span>
                                <a class="text-primary text-uppercase text-decoration-none" href="">Tours & Travel</a> -->
                                </div>
                                <router-link class="h5 m-0 text-decoration-none" to="/restaurantes">
                                Restaurante Cangrejo Azul
                                </router-link>
                            </div>
                        </div>
                    </div>
                    <div class="col-lg-4 col-md-6 mb-4 pb-2">
                        <div class="blog-item">
                            <div class="position-relative">
                                <img class="img-fluid w-100"
                                style="width:500px;height:250px;"   src="https://i.pinimg.com/736x/20/ff/08/20ff08999c08218b3dc5e92d8b06c61b.jpg"
                                    alt="">
                                <div class="blog-date">
                                    <h6 class="font-weight-bold mb-n1">Costa</h6>
                                    <small class="text-white text-uppercase">Brisa</small>
                                </div>
                            </div>
                            <div class="bg-white p-4">
                                <div class="d-flex mb-2">
                                    <!-- <a class="text-primary text-uppercase text-decoration-none" href="">Admin</a>
                                <span class="text-primary px-2">|</span>
                                <a class="text-primary text-uppercase text-decoration-none" href="">Tours & Travel</a> -->
                                </div>
                                <router-link class="h5 m-0 text-decoration-none" to="/restaurantes">
                                Restaurante Cangrejo Azul
                                </router-link>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <!-- Registration Start -->
        <div class="container-fluid bg-registration py-5" style="margin: 30px 0;">
            <div class="container py-5">
                <div class="row align-items-center">
                    <div class="col-lg-7 mb-5 mb-lg-0">
                        <div class="mb-4">
                            <h6 class="text-primary text-uppercase" style="letter-spacing: 5px;">Costa Brisa</h6>
                            <h1 class="text-white"><span class="text-primary">Ofrece</span></h1>
                        </div>
                        <p class="text-white">Dedica tus mañanas a relajarte en las hermosas playas.
                            Puedes disfrutar del sol, leer un buen libro o simplemente descansar bajo una sombrilla.
                            La <b>Costa Caribe,</b> es conocida por sus majestuosas playas, así que aprovecha al máximo este paraíso
                            tropical.</p>
                        <ul class="list-inline text-white m-0">
                            <li class="py-2"><i class="fa fa-check text-primary mr-3"></i>Playas</li>
                            <li class="py-2"><i class="fa fa-check text-primary mr-3"></i>Vista al Mar</li>
                            <li class="py-2"><i class="fa fa-check text-primary mr-3"></i>Recreaciones</li>
                        </ul>
                    </div>
                    <div class="col-lg-5">
                        <div style="border-radius: 10%; border-color: blue;" class="card border-0">

                            <img width="100%" height="300px" style="border-radius: 10%;"
                                src="https://img.freepik.com/foto-gratis/pareja-joven-mirando-camara-mientras-esta-sentado-tabla-surf-playa-sol_107420-9962.jpg?t=st=1710038492~exp=1710042092~hmac=d672b94ce166a1fb96a951464128a86c5f2b5ccad2b55f5c077f17a9436a6dc1&w=900"
                                alt="">

                        </div>
                    </div>
                </div>
            </div>
        </div>
        <!-- Registration End -->


        <!-- Team Start -->
        <!--    <div class="container-fluid py-6">
        <div class="container pt-6 pb-3">
            <div class="text-center mb-3 pb-3">
                <h6 class="text-primary text-uppercase" style="letter-spacing: 5px;">Guides</h6>
                <h1>Our Travel Guides</h1>
            </div>
            <div class="row">
                <div class="col-lg-3 col-md-4 col-sm-6 pb-2">
                    <div class="team-item bg-white mb-4">
                        <div class="team-img position-relative overflow-hidden">
                            <img class="img-fluid w-100" src="img/team-1.jpg" alt="">
                            <div class="team-social">
                                <a class="btn btn-outline-primary btn-square" href=""><i class="fab fa-twitter"></i></a>
                                <a class="btn btn-outline-primary btn-square" href=""><i class="fab fa-facebook-f"></i></a>
                                <a class="btn btn-outline-primary btn-square" href=""><i class="fab fa-instagram"></i></a>
                                <a class="btn btn-outline-primary btn-square" href=""><i class="fab fa-linkedin-in"></i></a>
                            </div>
                        </div>
                        <div class="text-center py-4">
                            <h5 class="text-truncate">Guide Name</h5>
                            <p class="m-0">Designation</p>
                        </div>
                    </div>
                </div>
                <div class="col-lg-3 col-md-4 col-sm-6 pb-2">
                    <div class="team-item bg-white mb-4">
                        <div class="team-img position-relative overflow-hidden">
                            <img class="img-fluid w-100" src="img/team-2.jpg" alt="">
                            <div class="team-social">
                                <a class="btn btn-outline-primary btn-square" href=""><i class="fab fa-twitter"></i></a>
                                <a class="btn btn-outline-primary btn-square" href=""><i class="fab fa-facebook-f"></i></a>
                                <a class="btn btn-outline-primary btn-square" href=""><i class="fab fa-instagram"></i></a>
                                <a class="btn btn-outline-primary btn-square" href=""><i class="fab fa-linkedin-in"></i></a>
                            </div>
                        </div>
                        <div class="text-center py-4">
                            <h5 class="text-truncate">Guide Name</h5>
                            <p class="m-0">Designation</p>
                        </div>
                    </div>
                </div>
                <div class="col-lg-3 col-md-4 col-sm-6 pb-2">
                    <div class="team-item bg-white mb-4">
                        <div class="team-img position-relative overflow-hidden">
                            <img class="img-fluid w-100" src="img/team-3.jpg" alt="">
                            <div class="team-social">
                                <a class="btn btn-outline-primary btn-square" href=""><i class="fab fa-twitter"></i></a>
                                <a class="btn btn-outline-primary btn-square" href=""><i class="fab fa-facebook-f"></i></a>
                                <a class="btn btn-outline-primary btn-square" href=""><i class="fab fa-instagram"></i></a>
                                <a class="btn btn-outline-primary btn-square" href=""><i class="fab fa-linkedin-in"></i></a>
                            </div>
                        </div>
                        <div class="text-center py-4">
                            <h5 class="text-truncate">Guide Name</h5>
                            <p class="m-0">Designation</p>
                        </div>
                    </div>
                </div>
                <div class="col-lg-3 col-md-4 col-sm-6 pb-2">
                    <div class="team-item bg-white mb-4">
                        <div class="team-img position-relative overflow-hidden">
                            <img class="img-fluid w-100" src="img/team-4.jpg" alt="">
                            <div class="team-social">
                                <a class="btn btn-outline-primary btn-square" href=""><i class="fab fa-twitter"></i></a>
                                <a class="btn btn-outline-primary btn-square" href=""><i class="fab fa-facebook-f"></i></a>
                                <a class="btn btn-outline-primary btn-square" href=""><i class="fab fa-instagram"></i></a>
                                <a class="btn btn-outline-primary btn-square" href=""><i class="fab fa-linkedin-in"></i></a>
                            </div>
                        </div>
                        <div class="text-center py-4">
                            <h5 class="text-truncate">Guide Name</h5>
                            <p class="m-0">Designation</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div> -->
        <!-- Team End -->


        <!-- Testimonial Start 
    <div class="container-fluid py-5">
        <div class="container py-5">
            <div class="text-center mb-3 pb-3">
                <h6 class="text-primary text-uppercase" style="letter-spacing: 5px;">Testimonial</h6>
                <h1>What Say Our Clients</h1>
            </div>
            <div class="owl-carousel testimonial-carousel">
                <div class="text-center pb-4">
                    <img class="img-fluid mx-auto" src="img/testimonial-1.jpg" style="width: 100px; height: 100px;" >
                    <div class="testimonial-text bg-white p-4 mt-n5">
                        <p class="mt-5">Dolor et eos labore, stet justo sed est sed. Diam sed sed dolor stet amet eirmod eos labore diam
                        </p>
                        <h5 class="text-truncate">Client Name</h5>
                        <span>Profession</span>
                    </div>
                </div>
                <div class="text-center">
                    <img class="img-fluid mx-auto" src="img/testimonial-2.jpg" style="width: 100px; height: 100px;" >
                    <div class="testimonial-text bg-white p-4 mt-n5">
                        <p class="mt-5">Dolor et eos labore, stet justo sed est sed. Diam sed sed dolor stet amet eirmod eos labore diam
                        </p>
                        <h5 class="text-truncate">Client Name</h5>
                        <span>Profession</span>
                    </div>
                </div>
                <div class="text-center">
                    <img class="img-fluid mx-auto" src="img/testimonial-3.jpg" style="width: 100px; height: 100px;" >
                    <div class="testimonial-text bg-white p-4 mt-n5">
                        <p class="mt-5">Dolor et eos labore, stet justo sed est sed. Diam sed sed dolor stet amet eirmod eos labore diam
                        </p>
                        <h5 class="text-truncate">Client Name</h5>
                        <span>Profession</span>
                    </div>
                </div>
                <div class="text-center">
                    <img class="img-fluid mx-auto" src="img/testimonial-4.jpg" style="width: 100px; height: 100px;" >
                    <div class="testimonial-text bg-white p-4 mt-n5">
                        <p class="mt-5">Dolor et eos labore, stet justo sed est sed. Diam sed sed dolor stet amet eirmod eos labore diam
                        </p>
                        <h5 class="text-truncate">Client Name</h5>
                        <span>Profession</span>
                    </div>
                </div>
            </div>
        </div>
    </div>
    Testimonial End -->
        <!-- Blog Start -->
        
        <!-- Blog End -->




        <HelloWorld />
    </div>


</template>

<script>
// @ is an alias to /src
import { watchEffect } from 'vue'
import HelloWorld from '@/components/HelloWorld.vue'
import { useAuth0 } from '@auth0/auth0-vue';
import axios from 'axios';

export default {
    name: 'HomeView',
    components: {
        HelloWorld
    },
    setup() {
        const { user, isAuthenticated } = useAuth0();
        const session = localStorage.getItem('usuarioAutenticado');
        const enviarUser = async (emaili) => {


            if (emaili !== undefined && emaili !== null && emaili !== "") {
                const response = await axios.post(`${process.env.API}/loginGoogle`, {
                    email: emaili,
                })

                if (response.data) {

                    if (response.data.token) {
                        localStorage.setItem('usuarioAutenticado', true);
                        localStorage.setItem('usuarioRol', response.data.user.rol);
                        localStorage.setItem('customerId', response.data.user._id);
                        localStorage.setItem('nombreUsuario', response.data.user.name);
                        localStorage.setItem('correoUsuario', response.data.user.email);
                        localStorage.setItem('celularUsuario', response.data.user.phone);
                        localStorage.setItem('imagenUsuario', response.data.user.image);
                        if (response.data.user.tipo_negocio !== undefined){
                            localStorage.setItem('Tipo_negocio', response.data.user.tipo_negocio);
                        }
                        if (localStorage.getItem('usuarioRol', response.data.user.rol) === 'Admin' || localStorage.getItem('usuarioRol', response.data.user.rol) === 'Cliente' || localStorage.getItem('usuarioRol', response.data.user.rol) === 'Vendedor') {
                            location.href = '/';
                            console.log("sesion iniciada")
                        } else {
                            this.$router.push({ path: '/login' });
                        }
                    }
                }
            } else {
                console.log("no hay datos")
            }

        };
        watchEffect(() => {

            if (!session && isAuthenticated.value && user.value.email) {
                enviarUser(user.value.email);
            }

        })


    },
    created() {

        /*    const { user,isAuthenticated } = useAuth0();
           
       if (isAuthenticated && user ) {
           
           console.log('HomeView setup', JSON.stringify(user.value, null, 2))
       } */
        /* this.$router.replace('/login') */


    }
}

</script>

<style>
.package-item img {
    width: 100%;
    height: 200px;
    /* Puedes ajustar la altura según tus preferencias */
    object-fit: cover;
    /* Ajusta la imagen manteniendo las proporciones y cubriendo completamente el contenedor */
}

/* Animación de caída desde la parte superior */
@keyframes fallIn {
    from {
        transform: translateY(-50%);
        opacity: 0;
    }

    to {
        transform: translateY(0);
        opacity: 1;
    }
}

/* Aplicar la animación a los elementos deseados */
.carousel-caption h4,
.carousel-caption h1 {
    animation: fallIn 1s ease-out;
    /* Puedes ajustar la duración y la función de temporización según tus preferencias */
}


/* Animación de pulsación para el botón "Solicitar Servicio" */
@keyframes pulse {
    0% {
        transform: scale(1);
    }

    50% {
        transform: scale(1.1);
    }

    100% {
        transform: scale(1);
    }
}

/* Aplicar la animación al botón "Solicitar Servicio" */
.btn-book-now {
    animation: pulse 1s infinite;
    /* Puedes ajustar la duración y la función de temporización según tus preferencias */
}





/* Ajuste de tamaño para las imágenes en el carrusel */
.carousel-inner img {
    height: 600px;
    /* Establece la altura según tus necesidades */
    width: 100%;
    /* Ajusta el ancho al 100% para mantener la proporción */
    object-fit: cover;
}

/* Aplicar sombreado azul al pasar el cursor sobre las imágenes */
.package-item img {
    transition: filter 0.3s ease-in-out;
}

.package-item:hover img {
    filter: brightness(80%) drop-shadow(0 0 10px rgba(0, 123, 255, 0.7));
}

/* #video-background {
  position: fixed;
  right: 0;
  bottom: 0;
  min-width: 100%;
  min-height: 100%;
  width: auto;
  height: auto;
  z-index: -100;
  background-size: cover;
} */

/* destinos  */
</style>
