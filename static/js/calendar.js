function deleteCalendarEvent(id) {
    $.ajax({
        url : `/calendar/${id}`,
        method : 'delete',
        success: function() {
            window.location.reload()
        }
    })
}