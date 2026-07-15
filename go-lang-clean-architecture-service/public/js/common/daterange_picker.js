function initDaterangePicker() {
  const $picker = $("#filter_daterange_picker");
  if (!$picker.length) return;

  /* 
    Check if URL has date params. If NOT, force "All Time" by clearing inputs.
    This overrides any default (like 30 days) set by server templates in the hidden inputs. 
  */
  const urlParams = new URLSearchParams(window.location.search);
  if (!urlParams.has("start_date") || !urlParams.has("end_date")) {
    $("#start_date_hidden").val("");
    $("#end_date_hidden").val("");
  }

  const startVal = $("#start_date_hidden").val();
  const endVal = $("#end_date_hidden").val();

  const start = startVal ? moment(startVal) : moment("2026-01-01");
  const end = endVal ? moment(endVal) : moment();

  createDATERANGEPicker(
    $picker,
    start,
    end,
    "#display-daterange",
    "#start_date_hidden",
    "#end_date_hidden",
    function () {
      htmx.trigger("body", "reloadPostList");
    }
  );
}

function createDATERANGEPicker(
  $el,
  start,
  end,
  displayId,
  startInputId,
  endInputId,
  callback
) {
  function updateDisplay(start, end, label) {
    if (label === "Tất cả thời gian") {
      $(displayId).text("Tất cả thời gian");
      $(startInputId).val("");
      $(endInputId).val("");
    } else {
      $(displayId).text(
        start.format("DD/MM/YYYY") + " - " + end.format("DD/MM/YYYY")
      );
      $(startInputId).val(start.format("YYYY-MM-DD"));
      $(endInputId).val(end.format("YYYY-MM-DD"));
    }
  }

  $el.daterangepicker(
    {
      startDate: start,
      endDate: end,
      ranges: {
        "Tất cả thời gian": [moment("2020-01-01"), moment()],
        "Hôm nay": [moment(), moment()],
        "Hôm qua": [moment().subtract(1, "days"), moment().subtract(1, "days")],
        "7 ngày qua": [moment().subtract(6, "days"), moment()],
        "30 ngày qua": [moment().subtract(29, "days"), moment()],
        "Tháng này": [moment().startOf("month"), moment().endOf("month")],
        "Tháng trước": [
          moment().subtract(1, "month").startOf("month"),
          moment().subtract(1, "month").endOf("month"),
        ],
      },
      locale: {
        format: "DD/MM/YYYY",
        applyLabel: "Áp dụng",
        cancelLabel: "Hủy",
        fromLabel: "Từ",
        toLabel: "Đến",
        customRangeLabel: "Tùy chọn",
        daysOfWeek: ["CN", "T2", "T3", "T4", "T5", "T6", "T7"],
        monthNames: [
          "Tháng 1",
          "Tháng 2",
          "Tháng 3",
          "Tháng 4",
          "Tháng 5",
          "Tháng 6",
          "Tháng 7",
          "Tháng 8",
          "Tháng 9",
          "Tháng 10",
          "Tháng 11",
          "Tháng 12",
        ],
        firstDay: 1,
      },
      opens: "left",
      alwaysShowCalendars: true,
    },
    function (start, end, label) {
      updateDisplay(start, end, label);
      if (callback) callback(start, end, label);
    }
  );

  if (!$(startInputId).val()) {
    updateDisplay(start, end, "Tất cả thời gian");
  } else {
    updateDisplay(start, end);
  }
}

export {initDaterangePicker}