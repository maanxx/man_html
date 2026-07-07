const blogMain = function () {
  function test() {
    console.log(123);
  }
  function test2() {
    console.log(456);
  }
  const tabs = document.querySelectorAll(".tabs-header_item");
  //   var tabsLength = tabs.length;

  //   for (var i = 0; i < tabsLength; i++) {
  //     tabs[i].addEventListener("click", function(){
  //         for (var j = 0; j< tabsLength; j++) {
  //             tabs[j].classList.remove("active")
  //         }
  //         this.classList.add("active");
  //     })
  //   }
  const tabContent = document.querySelectorAll(".tabs-content_item");

  tabs.forEach((tab) =>
    tab.addEventListener("click", function () {
      tabs.forEach((item) => item.classList.remove("active"));
      this.classList.add("active");
      var tabId = this.dataset.tab;
      tabContent.forEach((item) => {
        item.classList.remove("active");
      });
      document.getElementById(tabId).classList.add("active");
      //   var tabName = this.getAttribute("data-tab");
      //   alert(tabName);
      //   console.log(tabName);
    }),
  );

  return {
    init: function () {
      test();
      test2();
    },
  };
};

document.addEventListener("DOMContentLoaded", function (event) {
  blogMain().init();
});
