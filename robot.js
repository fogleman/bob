function roundedCylinder2(options) {
    options = options || {};
    var height = options.height || 1;
    var radius = options.radius || 1;
    var capRadius = options.capRadius || 0.1;
    var a = torus({
        ri: capRadius,
        ro: radius - capRadius
    }).translate([0, 0, capRadius]);
    var b = torus({
        ri: capRadius,
        ro: radius - capRadius
    }).translate([0, 0, height - capRadius]);
    var c = CSG.cylinder({
        start: [0, 0, capRadius],
        end: [0, 0, height - capRadius],
        radius: radius,
    });
    var d = CSG.cylinder({
        start: [0, 0, 0],
        end: [0, 0, height],
        radius: radius - capRadius,
    });
    return union(a, b, c, d);
}

function antenna() {
    return union(
        sphere(0.1).translate([0, 0, 0.5]),
        CSG.cylinder({
            start: [0, 0, -0.5],
            end: [0, 0, 0.5],
            radius: 0.05,
        })
    ).setColor([1, 1, 1]);
}

function head() {
    function headTop() {
        return difference(
            CSG.sphere({
                radius: 1,
                resolution: 32,
            }),
            CSG.cube({
                center: [0, 0, -10],
                radius: [10, 10, 10],
            })
        );
    }
    function headBottom() {
        return roundedCylinder2({
            height: 1.25,
            radius: 1,
            capRadius: 0.25,
        });
    }
    return union(
        headTop().translate([0, 0, 1]),
        headBottom()
    ).setColor([0.5, 0.5, 0.5]);
}

function eye() {
    return roundedCylinder2({
        height: 1,
        radius: 0.25,
        capRadius: 0.125/2,
    }).rotateX(90).translate([0, 0.5, 0]).setColor([1, 1, 1]);
}

function pupil() {
    return CSG.cylinder({
        start: [0, 0, 0],
        end: [0, 0, 1],
        radius: 0.15,
    }).rotateX(90).setColor([0, 0, 0]);
}

function neck() {
    return CSG.cylinder({
        start: [0, 0, 0],
        end: [0, 0, 1],
        radius: 0.5,
    }).setColor([1, 1, 1]);
}

function arm() {
    var end = roundedCylinder2({
        height: 0.25,
        radius: 0.375,
        capRadius: 0.125/2,
    }).translate([0, 0, -0.125]);
    var a = end.translate([0, 1, 0]);
    var b = end.translate([0, -1, 0]);
    var c = CSG.roundedCube({
        center: [0, 0, 0],
        radius: [0.375, 1+0.125/2, 0.125],
        roundradius: 0.125/2,
        resolution: 16,
    });
    return union(a, b, c).rotateY(90).translate([0, 1, 0]).setColor([1, 1, 1]);
}

function body() {
    var height = 3;
    var radius = 1;
    var capRadius = 0.125;
    var a = torus({
        ri: capRadius,
        ro: radius - capRadius
    }).translate([0, 0, capRadius]);
    var b = torus({
        ri: capRadius,
        ro: radius - capRadius
    }).translate([0, 0, height - capRadius]);
    var c = CSG.cylinder({
        start: [0, 0, capRadius],
        end: [0, 0, height - capRadius],
        radius: radius,
    });
    var d = CSG.cylinder({
        start: [0, 0, height / 2],
        end: [0, 0, height],
        radius: radius - capRadius,
    });
    return union(a, b, c, d).setColor([0.5, 0.5, 0.5]);
}

function wheel() {
    return CSG.sphere({
        radius: 0.75,
        resolution: 36,
    }).translate([0, 0, 0.75]).setColor([1, 1, 1]);
}

function code() {
    var spheres = [];
    for (var i = 0; i < 8; i++) {
        spheres.push(sphere(0.1).translate([0, 0, i * 0.25]));
    }
    return union(spheres).setColor([1, 1, 1]);
}

function robot() {
    return union(
        code().translate([0, 1-0.05, 1.5]),
        antenna().translate([0, 0, 6]),
        head().translate([0, 0, 4]),
        eye().translate([-0.4, 0.6, 4.75]),
        eye().translate([0.4, 0.6, 4.75]),
        pupil().scale([0.5, 0.5, 0.5]).translate([-0.4, 0.61+0.5, 4.75]),
        pupil().scale([0.5, 0.5, 0.5]).translate([0.4, 0.61+0.5, 4.75]),
        neck().translate([0, 0, 3.75]),
        body().translate([0, 0, 0.75]),
        arm().translate([0, 0, 0]).rotateX(-90).translate([-1.125, 0, 3]),
        arm().translate([0, 0, 0]).rotateX(0).translate([1.125, 0, 3]),
        wheel()
    );
}

function main() {
    return robot();
    return wheel();
    return arm();
    return body();
    return neck();
    return pupil();
    return eye();
    return head();
    return antenna();
    return code();
}
