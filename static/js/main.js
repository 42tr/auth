

(function ($) {
    "use strict";

    // Selectors
    const $validateForm = $('.validate-form');
    const $validateInputs = $('.validate-input .input100');

    // Function to display validation error
    function showValidation(inputElement) {
        $(inputElement).parent().addClass('alert-validate');
    }

    // Function to hide validation error
    function hideValidation(inputElement) {
        $(inputElement).parent().removeClass('alert-validate');
    }

    // Validation logic for an individual input
    function validateInput(inputElement) {
        const $input = $(inputElement);
        const inputType = $input.attr('type');
        const inputName = $input.attr('name');
        const inputValue = $input.val().trim();

        if (inputType === 'email' || inputName === 'email') {
            // Basic email validation regex
            const emailRegex = /^([a-zA-Z0-9_\-\.])@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([a-zA-Z0-9\-]+\.)+))([a-zA-Z]{1,5}|[0-9]{1,3})(\]?)$/;
            return emailRegex.test(inputValue);
        } else {
            // Check if field is not empty
            return inputValue !== '';
        }
    }

    // Handle form submission
    $validateForm.on('submit', function (event) {
        let formIsValid = true;

        $validateInputs.each(function () {
            if (!validateInput(this)) {
                showValidation(this);
                formIsValid = false;
            }
        });

        // Prevent form submission if validation fails
        if (!formIsValid) {
            event.preventDefault();
        }

        return formIsValid;
    });

    // Hide validation on input focus
    $validateInputs.on('focus', function () {
        hideValidation(this);
    });

})(jQuery);
