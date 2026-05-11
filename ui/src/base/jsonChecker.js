window.app = window.app || {};
window.app.components = window.app.components || {};

window.app.components.jsonChecker = function(json) {
    const valid = () => isValidStringifiedJSON(json() || "");

    return t.span(
        {
            className: "json-state",
            ariaDescription: app.attrs.tooltip(() => valid() ? "Valid JSON" : "Invalid JSON", "left"),
        },
        t.i({
            className: () => valid() ? "ri-checkbox-circle-fill txt-success" : "ri-error-warning-fill txt-danger",
            ariaHidden: true,
        }),
    );
};

function isValidStringifiedJSON(val) {
    console.log(val);
    if (val === "") {
        return true;
    }

    try {
        JSON.parse(val);
        return true;
    } catch {
        return false;
    }
}
