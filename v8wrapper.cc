#include <v8.h>
#include <string.h>

#include "v8wrapper.h"


char * runv8(const char *jssrc)
{

    v8::Isolate* isolate = v8::Isolate::New();

    v8::Handle<v8::Value> result;
    {
        v8::Isolate::Scope isolate_scope(isolate);

        // Create a stack-allocated handle scope.
        v8::HandleScope handle_scope(isolate);

        v8::Handle<v8::ObjectTemplate> global = v8::ObjectTemplate::New();

        v8::Handle<v8::Context> context = v8::Context::New(isolate, NULL, global);

        // Enter the created context for compiling and
        // running the script.
        v8::Context::Scope context_scope(context);

        // Create a string containing the JavaScript source code.
        v8::Handle<v8::String> source = v8::String::New(jssrc);

        // Compile the source code.
        v8::Handle<v8::Script> script = v8::Script::Compile(source);

        // Run the script
        result = script->Run();
    }

    // return result as string, must be deallocated in cgo wrapper
    v8::String::AsciiValue ascii(result);
    return strdup(*ascii);

}
