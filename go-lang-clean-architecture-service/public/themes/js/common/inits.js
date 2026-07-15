// const
const urlHome = "/cms";
const pathDefault = "images/default/";
const domain = window.location.origin;
const listStatusInit = {
  1: "ACTIVE",
  2: "INACTIVE",
  3: "DELETE",
  4: "ARCHIVE",
};

let accessToken = "";
let expireIn = 0;
let lastName = "";
let firstName = "";
let roleName = "";
let image = "";

function checkToken() {
  if (accessToken == "") {
    accessToken = localStorage.getItem("access_token");
  }
}
checkToken();
// const instanceAxios = axios.create({
//     baseURL: urlHome,
//     headers: {
//       Authorization: `Bearer `+ accessToken,
//     }
// });

const saveAuth = function saveAuth(data) {
  accessToken = data.access_token.token;
  expireIn = data.access_token.expire_in;
  lastName = data.user.last_name;
  firstName = data.user.first_name;
  image = data.user.image;
  roleName = data.user.roleName;

  localStorage.setItem("access_token", accessToken);
  localStorage.setItem("expire_in", expireIn);
  localStorage.setItem("first_name", firstName);
  localStorage.setItem("last_name", lastName);
  localStorage.setItem("image", image);
  localStorage.setItem("role_name", roleName);
};

const redirectURL = function redirectUrl(url = urlHome) {
  setTimeout(function () {
    window.location.href = url;
  }, 1200);
};

const showInfo = function showInfo() {
  lastName = localStorage.getItem("last_name");
  firstName = localStorage.getItem("first_name");
  image = localStorage.getItem("image");
  roleName = localStorage.getItem("role_name");
  $(".user-name-text").text(lastName + " " + firstName);
  $(".user-name-sub-text").text(roleName);
  $(".header-profile-user").attr("src", domain + image);
};

async function postAjax(url, paramObject) {}

function toName(str) {
  str = str.replace(/\s+/g, " ");
  str = str.trim();
  return str;
}

// replace title or slug
function toSlug(str) {
  str = str.toLowerCase();
  str = str.replace(/--\|-/g, "");
  str = str.replace(/^\-+|\-+$/g, "");
  str = str.replace(/à|á|ạ|ả|ã|â|ầ|ấ|ậ|ẩ|ẫ|ă|ằ|ắ|ặ|ẳ|ẵ/g, "a");
  str = str.replace(/è|é|ẹ|ẻ|ẽ|ê|ề|ế|ệ|ể|ễ/g, "e");
  str = str.replace(/ì|í|ị|ỉ|ĩ/g, "i");
  str = str.replace(/ò|ó|ọ|ỏ|õ|ô|ồ|ố|ộ|ổ|ỗ|ơ|ờ|ớ|ợ|ở|ỡ/g, "o");
  str = str.replace(/ù|ú|ụ|ủ|ũ|ư|ừ|ứ|ự|ử|ữ/g, "u");
  str = str.replace(/ỳ|ý|ỵ|ỷ|ỹ/g, "y");
  str = str.replace(/đ/g, "d");
  str = str.replace(
    /\`|\~|\!|\@|\#|\||\$|\%|\^|\&|\*|\(|\)|\+|\=|\,|\.|\/|\?|\>|\<|\'|\"|\:|\;|_/gi,
    "-",
  );
  str = str.replace(/[^a-zA-Z0-9 ]/g, "");
  str = str.replace(/-+/g, "-");
  str = str.replace(/^\-+|\-+$/g, "");
  str = str.replace(/\s+/g, "-");
  str = "@" + str + "@";
  str = str.replace(/\@\-|\-\@|\@/gi, "");
  return str;
}

function toSlugSimple(str) {
  str = str.toLowerCase();
  str = str.replace(/\s+/g, "-");
  str = str.replace(/à|á|ạ|ả|ã|â|ầ|ấ|ậ|ẩ|ẫ|ă|ằ|ắ|ặ|ẳ|ẵ/g, "a");
  str = str.replace(/è|é|ẹ|ẻ|ẽ|ê|ề|ế|ệ|ể|ễ/g, "e");
  str = str.replace(/ì|í|ị|ỉ|ĩ/g, "i");
  str = str.replace(/ò|ó|ọ|ỏ|õ|ô|ồ|ố|ộ|ổ|ỗ|ơ|ờ|ớ|ợ|ở|ỡ/g, "o");
  str = str.replace(/ù|ú|ụ|ủ|ũ|ư|ừ|ứ|ự|ử|ữ/g, "u");
  str = str.replace(/ỳ|ý|ỵ|ỷ|ỹ/g, "y");
  str = str.replace(/đ/g, "d");
  str = str.replace(/[`~!@#$%^&*()+=,./?><'":;_]/gi, "-");
  str = str.replace(/[^a-zA-Z0-9-]/g, "");
  str = str.replace(/-+/g, "-");
  str = str.replace(/^-+|-+$/g, "");
  return str;
}

function toSlugUppercase(str) {
  str = toSlug(str);
  str = str.toUpperCase();

  return str;
}

//max ordering for
function getMaxOrder(container, orderingClass) {
  var inputOrdering = container.find(orderingClass);
  var max = 0;
  inputOrdering.each(function (index, input) {
    var value = parseInt($(input).val());
    if (value > max) {
      max = value;
    }
  });
  return max + 1;
}

function showImage(input, img) {
  var reader = new FileReader();
  reader.addEventListener(
    "load",
    function () {
      img.src = reader.result;
    },
    false,
  );
  var file = input.files[0];
  if (file) {
    reader.readAsDataURL(file);
  }
}

//alert sweet success
function successAlert(msg) {
  Swal.fire({
    position: "top-end",
    icon: "success",
    title: msg,
    showConfirmButton: false,
    timer: 1500,
    showCloseButton: true,
  });
}

//alert sweet error
function errorAlert(msg) {
  Swal.fire({
    position: "top-end",
    title: msg,
    icon: "error",
    showConfirmButton: false,
    timer: 2500,
    showCloseButton: true,
  });
}

function forbiddenAlert() {
  Swal.fire({
    html:
      '<div class="mt-3"><lord-icon src="' +
      domain +
      '/static/themes/json/tdrtiskw.json" trigger="loop" colors="primary:#f06548,secondary:#f7b84b" style="width:120px;height:120px"></lord-icon><div class="mt-4 pt-2 fs-15"><h4>Oops...!</h4><p class="text-muted mx-4 mb-0">You do not have permission to perform this action</p></div></div>',
    showCancelButton: !0,
    showConfirmButton: !1,
    cancelButtonClass: "btn btn-primary w-xs mb-1",
    cancelButtonText: "Back",
    buttonsStyling: !1,
    showCloseButton: !0,
  });
}

function handleAjaxError(xhr) {
  if (xhr.status === 403) {
    forbiddenAlert();
    return;
  }
  const errResponse = JSON.parse(xhr.responseText);
  const errMessages = Array.isArray(errResponse.details.msg)
    ? errResponse.details.msg.join(" <br>")
    : errResponse.details.msg;
  errorAlert(errMessages);
}

function ischeckboxcheck() {
  Array.from(document.getElementsByName("chk_child")).forEach(function (a) {
    a.addEventListener("change", function (e) {
      1 == a.checked
        ? e.target.closest("tr").classList.add("table-active")
        : e.target.closest("tr").classList.remove("table-active");
      var t = document.querySelectorAll('[name="chk_child"]:checked').length;
      (e.target.closest("tr").classList.contains("table-active"),
        (document.getElementById("remove-actions").style.display =
          0 < t ? "block" : "none"));
    });
  });
  var checkboxesCount = document.querySelectorAll(
    '.form-check-all input[type="checkbox"]',
  ).length;
  var checkedCount = document.querySelectorAll(
    '.form-check-all input[type="checkbox"]:checked',
  ).length;
  checkboxesCount > 0 && checkboxesCount == checkedCount
    ? (checkAll.checked = true)
    : (checkAll.checked = false);
  document.getElementById("remove-actions").style.display =
    checkedCount > 0 ? "block" : "none";
}

function loader(isBool = true) {
  var preloader = document.getElementById("preloader");

  if (isBool) {
    preloader.style.opacity = "0.4";
    preloader.style.visibility = "visible";
  } else {
    preloader.style.opacity = "0";
    preloader.style.visibility = "hidden";
  }
}

function selectFileWithCKFinder(chooseEle, showEle) {
  CKFinder.modal({
    chooseFiles: true,
    width: 800,
    height: 600,
    onInit: function (finder) {
      finder.on("files:choose", function (evt) {
        var file = evt.data.files.first();
        var url = file.getUrl();
        url = url.replace(pathDefault, "");
        chooseEle.val(url);
        showEle.attr("src", url);
      });
      finder.on("file:choose:resizedImage", function (evt) {
        var output = chooseEle;
        output.val(evt.data.resizedUrl);
      });
    },
  });
}

function selectMultiFilesWithCKFinder(html, positionEle) {
  CKFinder.modal({
    chooseFiles: true,
    width: 800,
    height: 600,
    onInit: function (finder) {
      finder.on("files:choose", function (evt) {
        var files = evt.data.files;
        files.each(function (file, index) {
          finder
            .request("command:send", {
              name: "ImageInfo",
              folder: file.get("folder"),
              params: { fileName: file.get("name") },
            })
            .done(function (response) {
              if (response.error) {
                return;
              }
              var htmlStr = html.html();
              getLanguage();
              var url = file.getUrl();
              url = url.replace(pathDefault, "");

              htmlStr = htmlStr.replace(/{src}/g, url);
              htmlStr = htmlStr.replace(/{value_img}/g, url);
              htmlStr = htmlStr.replace(/{name}/g, file.attributes.name);
              htmlStr = htmlStr.replace(
                /{size}/g,
                parseInt(file.attributes.size),
              );
              htmlStr = htmlStr.replace(/{width}/g, response.width);
              htmlStr = htmlStr.replace(/{height}/g, response.height);

              positionEle.append(htmlStr);
            });
        });
      });
      finder.on("file:choose:resizedImage", function (evt) {
        var output = chooseEle;
        output.val(evt.data.resizedUrl);
      });
    },
  });
}

function showToolTip() {
  var tooltipTriggerList = [].slice.call(
    document.querySelectorAll('[data-bs-toggle="tooltip"]'),
  );
  tooltipTriggerList.map(function (tooltipTriggerEl) {
    var tooltip = new bootstrap.Tooltip(tooltipTriggerEl);

    tooltipTriggerEl.addEventListener("focusout", function () {
      tooltip.hide();
    });
    return tooltip;
  });
}

// Multi language setting
function getLanguage() {
  var language = localStorage.getItem("language");
  var request = new XMLHttpRequest();
  request.open("GET", "/static/themes/lang/" + language + ".json");
  request.onreadystatechange = function () {
    if (this.readyState === 4 && this.status === 200) {
      var data = JSON.parse(this.responseText);
      Object.keys(data).forEach(function (key) {
        var elements = document.querySelectorAll("[data-key='" + key + "']");
        Array.from(elements).forEach(function (elem) {
          if (elem.hasAttribute("data-bs-original-title")) {
            elem.setAttribute("data-bs-original-title", data[key]);
            return;
          }

          if (["INPUT", "TEXTAREA"].includes(elem.tagName)) {
            elem.setAttribute("placeholder", data[key]);
            return;
          }

          for (let i = 0; i < elem.childNodes.length; i++) {
            const node = elem.childNodes[i];
            if (node.nodeType === Node.TEXT_NODE) {
              node.textContent = data[key];
              break; // only change language first node text
            }
          }
        });
      });
    }
  };
  request.send();
}

function getLocalDateTime() {
  const now = new Date();
  const year = now.getFullYear();
  const month = String(now.getMonth() + 1).padStart(2, "0");
  const day = String(now.getDate()).padStart(2, "0");
  const hours = String(now.getHours()).padStart(2, "0");
  const minutes = String(now.getMinutes()).padStart(2, "0");
  const localDateTime = `${year}-${month}-${day} ${hours}:${minutes}`;
  return localDateTime;
}

function updatePagePath(slugId, pagePathId) {
  const slugInput = document.getElementById(slugId);
  const pagePath = document.getElementById(pagePathId);

  const pattern = pagePath.getAttribute("data-pattern");
  const slugValue = slugInput.value.trim();
  if (slugValue == "") {
    return (pagePath.textContent = pattern);
  }
  pagePath.textContent = pattern.replace("{slug}", slugValue);
}

function updateSlugIdPagePath(slugId, pagePathId, id) {
  const slugInput = document.getElementById(slugId);
  const pagePath = document.getElementById(pagePathId);
  if (!slugInput || !pagePath) return;

  const pattern = pagePath.getAttribute("data-pattern") || "";
  const slugValue = slugInput.value.trim();

  if (slugValue) {
    let result = pattern.replace("{slug}", slugValue);
    if (id) {
      result = result.replace("{id}", id);
    }
    pagePath.textContent = result;
  } else {
    pagePath.textContent = pattern;
  }
}

function updateURLParam(key, value) {
  const url = new URL(window.location);
  const isEmptyArray = Array.isArray(value) && value.length === 0;
  if (
    value === "all" ||
    value === undefined ||
    value === null ||
    value === "" ||
    isEmptyArray ||
    (key == "page" && value == 1)
  ) {
    url.searchParams.delete(key);
  } else {
    url.searchParams.set(key, encodeURIComponent(value));
  }

  window.history.pushState({}, "", url);
}

function decodeBase64Param(param) {
  try {
    return atob(decodeURIComponent(param));
  } catch (e) {
    return "";
  }
}

function autoExpandTextarea(root = document) {
  const textareas = root.querySelectorAll(".auto-expand");

  textareas.forEach((el) => {
    if (el.dataset.autoExpandInit === "true") return;

    const minHeight = parseInt(el.dataset.minHeight) || 80;

    const resize = () => {
      el.style.height = "auto";
      el.style.height = Math.max(el.scrollHeight, minHeight) + "px";
    };

    resize();
    el.addEventListener("input", resize);
    el.dataset.autoExpandInit = "true";
  });
}

function resetToDefaultText(root = document) {
  const textareas = root.querySelectorAll(".auto-expand");

  textareas.forEach((el) => {
    el.value = el.defaultValue;

    const minHeight = parseInt(el.dataset.minHeight) || 80;
    el.style.height = "auto";
    el.style.height = Math.max(el.scrollHeight, minHeight) + "px";

    el.dispatchEvent(new Event("input"));
  });
}

function showAutoExpandTextarea() {
  autoExpandTextarea();
  document.querySelectorAll(".modal").forEach((modal) => {
    modal.addEventListener("shown.bs.modal", () => {
      autoExpandTextarea(modal);
    });
  });
  document.querySelectorAll('[data-bs-toggle="pill"]').forEach((tab) => {
    tab.addEventListener("shown.bs.tab", (e) => {
      const target = document.querySelector(e.target.dataset.bsTarget);
      if (target) autoExpandTextarea(target);
    });
  });
}


export {
  accessToken,
  expireIn,
  lastName,
  firstName,
  pathDefault,
  listStatusInit,
  showInfo,
  saveAuth,
  redirectURL,
  toName,
  toSlug,
  toSlugUppercase,
  toSlugSimple,
  getMaxOrder,
  showImage,
  successAlert,
  errorAlert,
  handleAjaxError,
  ischeckboxcheck,
  loader,
  selectFileWithCKFinder,
  selectMultiFilesWithCKFinder,
  showToolTip,
  getLanguage,
  getLocalDateTime,
  forbiddenAlert,
  updatePagePath,
  updateSlugIdPagePath,
  updateURLParam,
  decodeBase64Param,
  showAutoExpandTextarea,
  resetToDefaultText
};
