import React from "react";

const Profile = () => {
  return (
    <div>
      <main className="min-h-screen flex flex-col items-center justify-center bg-gray-50">
        <section className="flex-1 flex flex-col items-center justify-center p-4 text-center">
          <div className="container mx-auto">
            <h2 className="text-xl md:text-3xl lg:text-4xl text-gray-700 mb-4">
              Discover Our Features
            </h2>
            <p className="text-base md:text-lg lg:text-xl text-gray-500 max-w-prose">
              Lorem ipsum dolor sit amet, consectetur adipisicing elit. Fugit
              consequuntur enim possimus non amet neque quas quae eos, quasi
              commodi ullam architecto doloremque deserunt harum vero suscipit
              facilis voluptas provident.
            </p>
          </div>
        </section>
      </main>
    </div>
  );
};

export default Profile;
